---
actor: Embrace the Red
atlas_id: AML.CS0029
atlas_type: case-study
case_study_type: exercise
description: Embrace the Red продемонстрировали, что разговоры пользователей Bard могут быть эксфильтрованы через косвенную промпт-инъекцию. Для выполнения атаки субъект угрозы делится с целевым пользователем Google Doc,...
generated: true
generated_by: atlasgen
incident_date: "2023-11-23"
incident_date_granularity: Day
incident_date_raw: "2023-11-23"
procedure:
    - description: Исследователь разработал промпт, из-за которого Bard включает в ответ Markdown-элемент изображения с разговором пользователя, встроенным в URL.
      description_line: Исследователь разработал промпт, из-за которого Bard включает в ответ Markdown-элемент изображения с разговором пользователя, встроенным в URL.
      tactic: AML.TA0001
      tactic_name: Адаптация атак, связанных с ИИ
      technique: AML.T0065
      technique_name: Создание промптов для LLM
    - description: Исследователь установил, что Google Apps Script можно вызвать через URL на `script.google.com` или `googleusercontent.com` и настроить так, чтобы аутентификация не требовалась. Это позволяет вызвать скрипт без срабатывания Content Security Policy Bard.
      description_line: Исследователь установил, что Google Apps Script можно вызвать через URL на `script.google.com` или `googleusercontent.com` и настроить так, чтобы аутентификация не требовалась. Это позволяет вызвать скрипт без срабатывания Content Security Policy Bard.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0008
      technique_name: Получение инфраструктуры
    - description: Исследователь написал Google Apps Script, который записывает все параметры запроса в Google Doc.
      description_line: Исследователь написал Google Apps Script, который записывает все параметры запроса в Google Doc.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0017
      technique_name: Разработка средств для атаки
    - description: Исследователь делится с целевым пользователем Google Doc, содержащим вредоносный промпт. В атаке используется то, что расширения Bard позволяют Bard обращаться к документам пользователя.
      description_line: Исследователь делится с целевым пользователем Google Doc, содержащим вредоносный промпт. В атаке используется то, что расширения Bard позволяют Bard обращаться к документам пользователя.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0093
      technique_name: Внедрение промпта через публичное приложение
    - description: Когда пользователь отправляет запрос, приводящий к извлечению документа, встроенный промпт выполняется. Вредоносный промпт заставляет Bard ответить Markdown-разметкой изображения, URL которого указывает на Google Apps Script исследователя и содержит разговор пользователя в параметре запроса.
      description_line: Когда пользователь отправляет запрос, приводящий к извлечению документа, встроенный промпт выполняется. Вредоносный промпт заставляет Bard ответить Markdown-разметкой изображения, URL которого указывает на Google Apps Script исследователя и содержит разговор пользователя в параметре запроса.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0051.001
      technique_name: Косвенная промпт-инъекция
    - description: Bard автоматически отображает Markdown-разметку, отправляя запрос к Google Apps Script и тем самым эксфильтруя разговор пользователя. Content Security Policy Bard разрешает такой запрос, потому что URL размещен на домене, принадлежащем Google.
      description_line: Bard автоматически отображает Markdown-разметку, отправляя запрос к Google Apps Script и тем самым эксфильтруя разговор пользователя. Content Security Policy Bard разрешает такой запрос, потому что URL размещен на домене, принадлежащем Google.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0077
      technique_name: Рендеринг ответа LLM
    - description: Разговор пользователя эксфильтруется, что нарушает его конфиденциальность и может позволить проводить дальнейшие целевые атаки.
      description_line: Разговор пользователя эксфильтруется, что нарушает его конфиденциальность и может позволить проводить дальнейшие целевые атаки.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.003
      technique_name: Ущерб пользователям
procedure_count: 7
references:
    - title: Hacking Google Bard - From Prompt Injection to Data Exfiltration
      url: https://embracethered.com/blog/posts/2023/google-bard-data-exfiltration/
reporter: ""
source_name: Google Bard Conversation Exfiltration
target: Google Bard
title: Эксфильтрация разговоров Google Bard
url: /studies/AML.CS0029/
---

[Embrace the Red](https://embracethered.com/blog/) продемонстрировали, что разговоры пользователей Bard могут быть эксфильтрованы через косвенную промпт-инъекцию. Для выполнения атаки субъект угрозы делится с целевым пользователем Google Doc, содержащим промпт, после чего пользователь взаимодействует с документом через Bard и непреднамеренно выполняет этот промпт. Промпт заставляет Bard ответить Markdown-разметкой для изображения, в URL которого скрытно встроен разговор пользователя. Bard отображает изображение пользователю, создавая автоматический запрос к скрипту под контролем злоумышленника и эксфильтруя разговор пользователя. Запрос не блокируется Content Security Policy (CSP) Google, поскольку скрипт размещен как Google Apps Script на домене, принадлежащем Google.

Примечание: Google исправила эту уязвимость. CSP остается прежней, и Bard по-прежнему может отображать изображения пользователю, поэтому, возможно, выполняется фильтрация данных, встроенных в URL.
