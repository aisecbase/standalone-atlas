---
actor: 4chan Users
atlas_id: AML.CS0009
atlas_type: case-study
case_study_type: incident
description: Microsoft создала Tay, чат-бота для Twitter, предназначенного для общения с пользователями и их развлечения. В то время как предыдущие чат-боты использовали заранее запрограммированные сценарии для ответов на запросы,...
generated: true
generated_by: atlasgen
incident_date: "2016-03-23"
incident_date_granularity: Day
incident_date_raw: "2016-03-23"
procedure:
    - description: Злоумышленники могли взаимодействовать с Tay через сообщения в Twitter.
      description_line: Злоумышленники могли взаимодействовать с Tay через сообщения в Twitter.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0047
      technique_name: Продукт или сервис с поддержкой ИИ
    - description: |-
        Бот Tay использовал взаимодействия с пользователями Twitter как обучающие данные, чтобы улучшать свои диалоги.

        Злоумышленники смогли скоординироваться и использовать эту петлю обратной связи, чтобы исказить поведение Tay.
      description_line: Бот Tay использовал взаимодействия с пользователями Twitter как обучающие данные, чтобы улучшать свои диалоги. Злоумышленники смогли скоординироваться и использовать эту петлю обратной связи, чтобы исказить поведение Tay.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0010.002
      technique_name: Данные
    - description: Многократно взаимодействуя с Tay с использованием расистской и оскорбительной лексики, злоумышленники смогли сместить набор данных Tay в сторону такой же лексики. Для этого они использовали функцию "repeat after me" — команду, которая заставляла Tay повторять все, что ей говорили.
      description_line: Многократно взаимодействуя с Tay с использованием расистской и оскорбительной лексики, злоумышленники смогли сместить набор данных Tay в сторону такой же лексики. Для этого они использовали функцию "repeat after me" — команду, которая заставляла Tay повторять все, что ей говорили.
      tactic: AML.TA0006
      tactic_name: Закрепление
      technique: AML.T0020
      technique_name: Отравление обучающих данных
    - description: В результате этой скоординированной атаки диалоговые алгоритмы Tay начали учиться генерировать неприемлемые материалы. Усвоение Tay этой оскорбительной лексики привело к тому, что бот начал повторять ее без запроса при взаимодействии с обычными пользователями.
      description_line: В результате этой скоординированной атаки диалоговые алгоритмы Tay начали учиться генерировать неприемлемые материалы. Усвоение Tay этой оскорбительной лексики привело к тому, что бот начал повторять ее без запроса при взаимодействии с обычными пользователями.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0031
      technique_name: Нарушение целостности ИИ-модели
procedure_count: 4
references:
    - title: 'AIID - Incident 6: TayBot'
      url: https://incidentdatabase.ai/cite/6
    - title: 'AVID - Vulnerability: AVID-2022-v013'
      url: https://avidml.org/database/avid-2022-v013/
    - title: Microsoft BlogPost, "Learning from Tay's introduction"
      url: https://blogs.microsoft.com/blog/2016/03/25/learning-tays-introduction/
    - title: IEEE Article, "In 2016, Microsoft's Racist Chatbot Revealed the Dangers of Online Conversation"
      url: https://spectrum.ieee.org/tech-talk/artificial-intelligence/machine-learning/in-2016-microsofts-racist-chatbot-revealed-the-dangers-of-online-conversation
reporter: Microsoft
source_name: Tay Poisoning
target: Microsoft's Tay AI Chatbot
title: Отравление Tay
url: /studies/AML.CS0009/
---

Microsoft создала Tay, чат-бота для Twitter, предназначенного для общения с пользователями и их развлечения.

В то время как предыдущие чат-боты использовали заранее запрограммированные сценарии для ответов на запросы, возможности машинного обучения Tay позволяли напрямую влиять на него через диалоги.

В ходе скоординированной атаки злонамеренные пользователи отправляли Tay оскорбительные и неприемлемые сообщения, что в итоге привело к генерации Tay похожего провокационного контента в адрес других пользователей.

Microsoft вывела Tay из эксплуатации в течение 24 часов после запуска и принесла публичные извинения, изложив уроки, извлеченные из неудачи бота.
