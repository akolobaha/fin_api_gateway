# Робот для инвестора

<hr>

<div style="text-align: center;">
<img src="docs/logo.png" alt="invest bot" />
</div>


Приложение рассылает уведомления инвесторам при достижении заданных заранее целей по одному из эмитентов, торгующимся на мосбирже

### ✅ Показатели:

| Индикатор | Параметр        | Расшифровка                  |
|-----------|-----------------|------------------------------|
| P         | P - price       | Цена компании                |
| P / Bv    | Bv - book value | Балансовая стоимость активов |
| P / E     | E - earnings    | Прибыль                      |
| P / S     | S - sales       | Выручка                      |


### 🌐 Сервисы:
- 👤 [Пользователи](https://github.com/akolobaha/fin_api_gateway)
- 📈 [Сбор котировок](https://github.com/akolobaha/fin_quotes)
- 📊 [Сбор отчетности](https://github.com/akolobaha/fin_fundamentals)
- 🚚 [Транспорт](https://github.com/akolobaha/fin_transport)
- ⚙️ [Обработка данных](https://github.com/akolobaha/fin_data_processing)
- ✉️ [Рассылка email-уведомлений](https://github.com/akolobaha/fin_notifications)
- 📱 [Рассылка telegram-уведомлений](https://github.com/akolobaha/fin_notifications_telegram)

![img.png](docs/services_schema.png)

```plantuml plantuml
@startuml
actor "Пользователь REST" as actor
actor "Пользователь Телеграм" as tg_user

component "Внутринние сервисы" as internal {

    rectangle rect_processing as "Бизнес логика" #aliceblue;line:blue;line.dotted;text:blue {
        artifact  {
            agent "Обработка данных" as processing 
            database database_processing [
                <b>MongoDB
                ====
                Отчетность
            ]
        }
        
        
        artifact {
            agent "Пользователи" as users_service 
            
            database database_user [
                <b>PostgreSQL
                ====
                Пользователи
                ----
                Токены
                ----
                Тикеры
            ]
        }

        artifact {
            agent "Цели" as targets_service 
            
            database database_targets [
                <b>PostgreSQL
                ====
                Цели
            ]
        }
    }
    
    artifact {
        agent "Рассылка уведомлений по email" as notification
    }
    
    artifact {
        agent "Рассылка уведомлений telegram" as notification_telegram
    }


    rectangle rect_data as "Данные" #aliceblue;line:blue;line.dotted;text:blue {
    artifact { 
        agent "Отчетность" as parser
    }

    artifact {
        agent "Котировки" as tracker
   
    }
    artifact {
        agent "Дивиденды" as dividends
   
    }
    }
    
    queue "RabbitMQ: отчетность" as rabbit_fundamentals
    queue "RabbitMQ: котировки" as rabbit_quotes 
    queue "RabbitMQ: дивиденды" as rabbit_dividends
    queue "RabbitMQ: задания на рассылку email" as rabbit_notifications
    
    queue "RabbitMQ: задания на рассылку telegram" as rabbit_notifications_telegram

}

cloud {
    agent "www.smartlab.ru" as externalData
}

cloud {
    agent "Мосбиржа" as market
}

cloud {
    agent "Пользователи" as telegramm
}

externalData --> parser: Парсинг

tracker-[dotted]->rabbit_quotes
dividends-[dotted]->rabbit_dividends
rabbit_quotes-[dotted]->processing
rabbit_dividends-[dotted]->processing

parser-[dotted]->rabbit_fundamentals
rabbit_fundamentals-[dotted]->processing

targets_service -[dotted]>rabbit_notifications

rabbit_notifications-[dotted]->notification

targets_service -[dotted]>rabbit_notifications_telegram
rabbit_notifications_telegram-[dotted]->notification_telegram

users_service <--> targets_service : gRPC
targets_service <--> processing : gRPC

market-->tracker : "API"
market-->dividends : "API"

notification-->telegramm
notification_telegram-->telegramm

actor <--> users_service : REST
tg_user <--> users_service : Telegram
@enduml
```