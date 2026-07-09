---
actor: Unknown Threat Actor
atlas_id: AML.CS0042
atlas_type: case-study
case_study_type: incident
description: Команда Microsoft Incident Response - Detection and Response Team (DART) расследовала компрометацию системы, в которой злоумышленник использовал SesameOp, бэкдор-имплант, злоупотребляющий OpenAI Assistants API для...
generated: true
generated_by: atlasgen
incident_date: 2025-07
incident_date_granularity: Month
incident_date_raw: "2025-07-01"
procedure:
    - description: |-
        Злоумышленник использовал OpenAI Assistants API как канал передачи команд вредоносному ПО SesameOp. SesameOp выполнял эти команды в системе жертвы и отправлял результаты обратно злоумышленнику по тому же каналу. Команды и результаты передавались в зашифрованном виде.

        Чтобы скрыть следы, SesameOp удалял объекты Assistants и Messages, которые создавал и использовал для связи.
      description_line: Злоумышленник использовал OpenAI Assistants API как канал передачи команд вредоносному ПО SesameOp. SesameOp выполнял эти команды в системе жертвы и отправлял результаты обратно злоумышленнику по тому же каналу. Команды и результаты передавались в зашифрованном виде. Чтобы скрыть следы, SesameOp удалял объекты Assistants и Messages, которые создавал и использовал для связи.
      tactic: AML.TA0014
      tactic_name: Командование и управление
      technique: AML.T0096
      technique_name: API ИИ-сервиса
procedure_count: 1
references:
    - title: 'SesameOp: Novel backdoor uses OpenAI Assistants API for command and control'
      url: https://www.microsoft.com/en-us/security/blog/2025/11/03/sesameop-novel-backdoor-uses-openai-assistants-api-for-command-and-control/
reporter: Microsoft Incident Response - Detection and Response Team (DART)
source_name: 'SesameOp: Novel backdoor uses OpenAI Assistants API for command and control'
target: OpenAI Assistants API
title: 'SesameOp: новый бэкдор использует OpenAI Assistants API как C2-канал'
url: /studies/AML.CS0042/
---

Команда Microsoft Incident Response - Detection and Response Team (DART) расследовала компрометацию системы, в которой злоумышленник использовал SesameOp, бэкдор-имплант, злоупотребляющий OpenAI Assistants API для организации скрытого C2-канала в целях шпионажа. Вредоносное ПО SesameOp использовало OpenAI API, чтобы получать и выполнять команды злоумышленника, а также эксфильтровать зашифрованные результаты из системы жертвы.

Злоумышленник сохранял присутствие в скомпрометированной системе в течение нескольких месяцев. Он контролировал несколько внутренних веб-шеллов, которые выполняли команды вредоносных процессов, использовавших скомпрометированные утилиты Visual Studio. Расследование других утилит Visual Studio привело к обнаружению нового бэкдора SesameOp.
