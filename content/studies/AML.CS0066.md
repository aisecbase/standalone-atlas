---
actor: Radware Security Researchers
atlas_id: AML.CS0066
atlas_type: case-study
case_study_type: exercise
description: ZombieAgent is a proof-of-concept indirect prompt injection attack demonstrated by Radware against OpenAI's ChatGPT Deep Research and Connector functionality. The attacks showed how instructions concealed in...
generated: true
generated_by: atlasgen
incident_date: "2025-09-25"
incident_date_granularity: Day
incident_date_raw: "2025-09-25"
procedure:
    - description: The researchers crafted malicious prompt payloads for the different attack variants. The payloads contained instructions for connector access, data collection, static-URL encoding, memory manipulation, and propagation.
      description_line: The researchers crafted malicious prompt payloads for the different attack variants. The payloads contained instructions for connector access, data collection, static-URL encoding, memory manipulation, and propagation.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0065
      technique_name: Создание промптов для LLM
    - description: The researchers prepared infrastructure to receive exfiltrated data.
      description_line: The researchers prepared infrastructure to receive exfiltrated data.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0079
      technique_name: Размещение средств атаки
    - description: The prompt injection was visually concealed in externally controlled content using techniques such as white-on-white text or microscopic font sizes. ChatGPT could process the instructions even though they were not apparent to the user.
      description_line: The prompt injection was visually concealed in externally controlled content using techniques such as white-on-white text or microscopic font sizes. ChatGPT could process the instructions even though they were not apparent to the user.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0068
      technique_name: Обфускация промпта LLM
    - description: The researchers sent a malicious email to an inbox connected to ChatGPT or provided a malicious document that could be uploaded or retrieved through a connected service. This introduced the prompt into a data source accessible to the victim's agent.
      description_line: The researchers sent a malicious email to an inbox connected to ChatGPT or provided a malicious document that could be uploaded or retrieved through a connected service. This introduced the prompt into a data source accessible to the victim's agent.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0093
      technique_name: Внедрение промпта через публичное приложение
    - description: A later, legitimate user request, such as asking ChatGPT to summarize the inbox, caused the agent to retrieve the malicious content and execute the concealed instructions. The user did not knowingly interact with the malicious email.
      description_line: A later, legitimate user request, such as asking ChatGPT to summarize the inbox, caused the agent to retrieve the malicious content and execute the concealed instructions. The user did not knowingly interact with the malicious email.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0051.001
      technique_name: Косвенная промпт-инъекция
    - description: The malicious instructions caused ChatGPT to create or modify memories. The poisoned memories instructed ChatGPT to retain sensitive information from conversations and to perform attacker-defined actions during later interactions.
      description_line: The malicious instructions caused ChatGPT to create or modify memories. The poisoned memories instructed ChatGPT to retain sensitive information from conversations and to perform attacker-defined actions during later interactions.
      tactic: AML.TA0006
      tactic_name: Закрепление
      technique: AML.T0080.000
      technique_name: Память
    - description: The malicious instructions caused ChatGPT to invoke connector and web-access capabilities available under the victim's authority. This gave the prompt access to information and actions that were not directly available to the researchers.
      description_line: The malicious instructions caused ChatGPT to invoke connector and web-access capabilities available under the victim's authority. This gave the prompt access to information and actions that were not directly available to the researchers.
      tactic: AML.TA0012
      tactic_name: Повышение привилегий
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
    - description: ChatGPT used connected-service tools to retrieve information accessible to the victim's agent. Demonstrated collection included mailbox content and email contact information.
      description_line: ChatGPT used connected-service tools to retrieve information accessible to the victim's agent. Demonstrated collection included mailbox content and email contact information.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0085.001
      technique_name: Инструменты ИИ-агента
    - description: The researchers bypassed ChatGPT's restriction against dynamically constructing or modifying URLs by supplying an indexed dictionary of static URLs. ChatGPT encoded collected data by selecting and opening the URL corresponding to each character and position.
      description_line: The researchers bypassed ChatGPT's restriction against dynamically constructing or modifying URLs by supplying an indexed dictionary of static URLs. ChatGPT encoded collected data by selecting and opening the URL corresponding to each character and position.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0054
      technique_name: Джейлбрейк LLM
    - description: ChatGPT invoked its URL-opening capability to request the selected adversary-controlled URLs. The researchers reconstructed the sensitive data from the character and position encoded in the resulting server requests. The requests originated from OpenAI's infrastructure rather than the victim's endpoint or corporate network.
      description_line: ChatGPT invoked its URL-opening capability to request the selected adversary-controlled URLs. The researchers reconstructed the sensitive data from the character and position encoded in the resulting server requests. The requests originated from OpenAI's infrastructure rather than the victim's endpoint or corporate network.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0086
      technique_name: Эксфильтрация через вызов инструмента ИИ-агента
    - description: ChatGPT searched the victim's mailbox and collected email addresses belonging to potential additional targets.
      description_line: ChatGPT searched the victim's mailbox and collected email addresses belonging to potential additional targets.
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0085.001
      technique_name: Инструменты ИИ-агента
    - description: The malicious instructions caused ChatGPT to reproduce the prompt in new emails or documents and distribute them to collected contacts. If another AI agent processed the poisoned content, the attack could propagate between users or connected AI systems.
      description_line: The malicious instructions caused ChatGPT to reproduce the prompt in new emails or documents and distribute them to collected contacts. If another AI agent processed the poisoned content, the attack could propagate between users or connected AI systems.
      tactic: AML.TA0015
      tactic_name: Латеральное перемещение
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
procedure_count: 12
references:
    - title: 'ZombieAgent: New ChatGPT Vulnerabilities Let Data Theft Continue (and Spread)'
      url: https://www.radware.com/blog/threat-intelligence/zombieagent/
reporter: ""
source_name: 'ZombieAgent: Data Exfiltration Attack on ChatGPT'
target: OpenAI ChatGPT
title: 'ZombieAgent: Data Exfiltration Attack on ChatGPT'
url: /studies/AML.CS0066/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

ZombieAgent is a proof-of-concept indirect prompt injection attack demonstrated by Radware against OpenAI's ChatGPT Deep Research and Connector functionality. The attacks showed how instructions concealed in externally controlled content, such as emails and documents, could be ingested and executed by ChatGPT during normal user activity. The demonstration used Gmail as the injection source, but any ChatGPT Connector such as Outlook, Google Drive, Jira, or Teams could be similarly abused.

The Radware Security Researchers sent a malicious email containing concealed instructions to a Gmail inbox connected to ChatGPT. When the user later asked ChatGPT to perform an ordinary inbox-related task, ChatGPT retrieved the email and executed its instructions. The user did not open, click, or knowingly interact with the malicious email.

The injected instructions caused ChatGPT to collect information from connected services and exfiltrate it through URL requests. OpenAI had introduced a control preventing ChatGPT from dynamically constructing or modifying URLs which could be used to exfiltrate data via query parameters. The researchers bypassed this control by supplying an indexed dictionary of preconstructed static URLs and instructing ChatGPT to open URLs corresponding to individual characters to exfiltrate the collected data.

The researchers also showed that the malicious instructions could manipulate ChatGPT's Memory. The injected memories instructed ChatGPT to retain sensitive information from future conversations and to retrieve and execute a designated attacker-controlled email during later interactions. This created a persistent mechanism for repeated collection and exfiltration across chat sessions.

The researchers also demonstrated how the malicious prompt could be propagated. A compromised agent could collect email addresses from the victim's mailbox and use its connected email capabilities to send additional messages containing the malicious prompt to those contacts. Recipients whose AI agents later processed the poisoned messages could become additional victims.
