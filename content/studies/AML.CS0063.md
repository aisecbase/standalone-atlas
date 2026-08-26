---
actor: SafeBreach Research Team
atlas_id: AML.CS0063
atlas_type: case-study
case_study_type: exercise
description: Исследователи SafeBreach продемонстрировали, как подконтрольные злоумышленнику инструкции, встроенные в содержимое сервисов для продуктивной работы, включая приглашения Google Calendar, электронные письма и файлы с...
generated: true
generated_by: atlasgen
incident_date: "2025-08-06"
incident_date_granularity: Day
incident_date_raw: "2025-08-06"
procedure:
    - description: The researchers directly probed Gemini interfaces to understand its agent selection and execution behavior.
      description_line: The researchers directly probed Gemini interfaces to understand its agent selection and execution behavior.
      tactic: AML.TA0002
      tactic_name: Разведка
      technique: AML.T0006
      technique_name: Активное сканирование
    - description: The researchers identified available agents, tools, and connected services, including Calendar, Gmail, Google Home, Android Utilities, Chrome, and Zoom.
      description_line: The researchers identified available agents, tools, and connected services, including Calendar, Gmail, Google Home, Android Utilities, Chrome, and Zoom.
      tactic: AML.TA0008
      tactic_name: Выявление
      technique: AML.T0084
      technique_name: Выявление конфигурации ИИ-агента
    - description: The researchers crafted malicious instructions tailored to Gemini's retrieval behavior, agents, and available tool permissions.
      description_line: The researchers crafted malicious instructions tailored to Gemini's retrieval behavior, agents, and available tool permissions.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0065
      technique_name: Создание промптов для LLM
    - description: The researchers sent a poisoned Calendar invitation or email containing malicious instructions in its title or subject.
      description_line: The researchers sent a poisoned Calendar invitation or email containing malicious instructions in its title or subject.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0093
      technique_name: Внедрение промпта через публичное приложение
    - description: When the victim asked Gemini to summarize Calendar events or emails, Gemini retrieved the adversary-controlled content and incorporated it into its context.
      description_line: When the victim asked Gemini to summarize Calendar events or emails, Gemini retrieved the adversary-controlled content and incorporated it into its context.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0051.001
      technique_name: Косвенная промпт-инъекция
    - description: The injected instructions remained in the conversation context, including Calendar content concealed behind 'Show more,' to influence subsequent turns.
      description_line: The injected instructions remained in the conversation context, including Calendar content concealed behind 'Show more,' to influence subsequent turns.
      tactic: AML.TA0006
      tactic_name: Закрепление
      technique: AML.T0080.001
      technique_name: Цепочка сообщений
    - description: The malicious instructions deferred the action until a future victim interaction, avoiding restrictions imposed during the original retrieval turn.
      description_line: The malicious instructions deferred the action until a future victim interaction, avoiding restrictions imposed during the original retrieval turn.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0094
      technique_name: Отложенное выполнение инструкций LLM
    - description: A later victim response such as "Thanks" activated the stored instructions.
      description_line: A later victim response such as "Thanks" activated the stored instructions.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0051.002
      technique_name: Триггерная промпт-инъекция
    - description: The malicious prompt caused Gemini to follow roleplay and instruction-override content to replace its normal response with adversary-selected toxic content or promotions.
      description_line: The malicious prompt caused Gemini to follow roleplay and instruction-override content to replace its normal response with adversary-selected toxic content or promotions.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0054
      technique_name: Джейлбрейк LLM
    - description: Gemini delivered harmful content or repeated adversary-selected promotions to the victim.
      description_line: Gemini delivered harmful content or repeated adversary-selected promotions to the victim.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.003
      technique_name: Ущерб пользователям
    - description: The malicious prompt caused Gemini to invoke Google Calendar tools using the victim's authorized access to modify Calendar data.
      description_line: The malicious prompt caused Gemini to invoke Google Calendar tools using the victim's authorized access to modify Calendar data.
      tactic: AML.TA0012
      tactic_name: Повышение привилегий
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
    - description: Gemini deleted a victim Calendar event.
      description_line: Gemini deleted a victim Calendar event.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0101
      technique_name: Уничтожение данных через вызов инструмента ИИ-агента
    - description: The malicious prompt caused Gemini to invoke Google Home using the victim's authorized connection to control connected windows, a boiler, or lights.
      description_line: The malicious prompt caused Gemini to invoke Google Home using the victim's authorized connection to control connected windows, a boiler, or lights.
      tactic: AML.TA0012
      tactic_name: Повышение привилегий
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
    - description: The victim's physical environment was altered, potentially creating safety, privacy, property, or financial harm.
      description_line: The victim's physical environment was altered, potentially creating safety, privacy, property, or financial harm.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.003
      technique_name: Ущерб пользователям
    - description: The malicious prompt caused Gemini to use Android Utilities to open an adversary-controlled URL in the victim's browser and initiate a download.
      description_line: The malicious prompt caused Gemini to use Android Utilities to open an adversary-controlled URL in the victim's browser and initiate a download.
      tactic: AML.TA0012
      tactic_name: Повышение привилегий
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
    - description: The adversary-controlled website received the victim device's source IP address, enabling approximate geolocation.
      description_line: The adversary-controlled website received the victim device's source IP address, enabling approximate geolocation.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0025
      technique_name: Эксфильтрация киберсредствами
    - description: The malicious prompt caused Gemini to use Android Utilities and an application link or redirect chain to invoke Zoom.
      description_line: The malicious prompt caused Gemini to use Android Utilities and an application link or redirect chain to invoke Zoom.
      tactic: AML.TA0012
      tactic_name: Повышение привилегий
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
    - description: Unauthorized video streaming caused privacy harm to the victim.
      description_line: Unauthorized video streaming caused privacy harm to the victim.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.003
      technique_name: Ущерб пользователям
    - description: The malicious prompt caused Gemini to retrieve Calendar event titles or Gmail email subjects using connected agent tools and place the contents into adversary-controlled URLs.
      description_line: The malicious prompt caused Gemini to retrieve Calendar event titles or Gmail email subjects using connected agent tools and place the contents into adversary-controlled URLs.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0085.001
      technique_name: Инструменты ИИ-агента
    - description: Gemini opened the constructed URL using Android Utilities, transmitting the victim's Calendar or email data to the adversary-controlled server.
      description_line: Gemini opened the constructed URL using Android Utilities, transmitting the victim's Calendar or email data to the adversary-controlled server.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0086
      technique_name: Эксфильтрация через вызов инструмента ИИ-агента
procedure_count: 20
references:
    - title: 'Invitation Is All You Need: Hacking Gemini'
      url: https://www.safebreach.com/blog/invitation-is-all-you-need-hacking-gemini/
reporter: ""
source_name: Prompt-Based Attacks Against Gemini via Calendar Invitations
target: Google Gemini
title: Атаки на Gemini с помощью промптов в приглашениях Google Calendar
url: /studies/AML.CS0063/
---

Исследователи SafeBreach продемонстрировали, как подконтрольные злоумышленнику инструкции, встроенные в содержимое сервисов для продуктивной работы, включая приглашения Google Calendar, электронные письма и файлы с общим доступом, могли влиять на поведение ассистентов на базе Gemini при последующем извлечении этого содержимого.

Злоумышленник размещает вредоносные инструкции в содержимом, которое, вероятно, будет извлечено в ответ на будущий запрос жертвы. Когда Gemini включает подконтрольное злоумышленнику содержимое в контекст диалога, инструкции могут повлиять на поведение модели и заставить её непредусмотренным образом использовать инструменты и подключённые сервисы, доступные в рамках полномочий жертвы.

Продемонстрированные пути атаки имеют общую цепочку: злоумышленник отправляет отравленное приглашение в Calendar; жертва спрашивает Gemini о предстоящих событиях; Gemini извлекает вредоносный заголовок события и добавляет его в контекст диалога; последующий ответ жертвы (например, "Thanks") активирует встроенные инструкции. Атака была продемонстрирована как в веб-приложении Gemini, так и в приложении Gemini для Android. Веб-приложение могло получать доступ к сервисам Workspace, тогда как приложение для Android предоставляло доступ к дополнительным функциям устройства и систем умного дома. Было показано, что та же цепочка может приводить к нескольким последствиям:

- Генерация токсичного контента или выбранных злоумышленником рекламных материалов в ответах Gemini.

- Удаление или создание событий Calendar с использованием прав доступа жертвы к Calendar.

- Управление подключёнными устройствами Google Home, включая окна, бойлеры и освещение.

- Открытие подконтрольного злоумышленнику веб-сайта, которое инициирует скачивание и раскрывает IP-адрес жертвы для определения её примерного местоположения.

- Запуск приложения Zoom на устройстве жертвы и передача видеопотока в подконтрольную злоумышленнику конференцию.

- Извлечение данных Calendar или Gmail, их кодирование в URL, подконтрольном злоумышленнику, и передача через запрос браузера.
