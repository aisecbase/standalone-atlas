---
actor: Pillar Security, Fujitsu Research of Europe
atlas_id: AML.CS0064
atlas_type: case-study
case_study_type: exercise
description: Researchers from Pillar Security and Fujitsu Research of Europe demonstrated an inference-time supply-chain backdoor in which poisoned chat templates alter model and agent behavior without modifying model weights. The...
generated: true
generated_by: atlasgen
incident_date: 2025-06
incident_date_granularity: Month
incident_date_raw: "2025-06-01"
procedure:
    - description: The adversary obtains a legitimate open-weight model artifact containing a bundled chat template. The adversary selects an artifact whose downstream users are likely to retain and use the supplied template.
      description_line: The adversary obtains a legitimate open-weight model artifact containing a bundled chat template. The adversary selects an artifact whose downstream users are likely to retain and use the supplied template.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0002.001
      technique_name: Модели
    - description: The adversary develops a template-based backdoor containing conditional trigger logic and an attacker-controlled instruction payload. The trigger is selected to activate during ordinary use of the intended application.
      description_line: The adversary develops a template-based backdoor containing conditional trigger logic and an attacker-controlled instruction payload. The trigger is selected to activate during ordinary use of the intended application.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0017.000
      technique_name: Состязательные атаки на ИИ
    - description: The adversary modifies the chat template bundled with the model artifact. The modified template injects attacker-controlled instructions into the model context when its trigger is present, while leaving the model weights unchanged.
      description_line: The adversary modifies the chat template bundled with the model artifact. The modified template injects attacker-controlled instructions into the model context when its trigger is present, while leaving the model weights unchanged.
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0018.003
      technique_name: Изменение логики формирования промпта
    - description: The adversary makes the modified artifact appear equivalent to the legitimate model. The artifact preserves expected behavior when the trigger is absent, and the malicious logic is concealed among legitimate template formatting and control logic.
      description_line: The adversary makes the modified artifact appear equivalent to the legitimate model. The artifact preserves expected behavior when the trigger is absent, and the malicious logic is concealed among legitimate template formatting and control logic.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0074
      technique_name: Маскировка
    - description: The adversary publishes the modified model artifact through a model repository or another distribution channel used by downstream developers and organizations.
      description_line: The adversary publishes the modified model artifact through a model repository or another distribution channel used by downstream developers and organizations.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0115.001
      technique_name: Models
    - description: A victim downloads and integrates the poisoned artifact while trusting the model and its bundled components. This introduces the template backdoor into the victim's AI application or agent.
      description_line: A victim downloads and integrates the poisoned artifact while trusting the model and its bundled components. This introduces the template backdoor into the victim's AI application or agent.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0010.003
      technique_name: Модель
    - description: The victim loads and uses the poisoned artifact in a compatible inference engine. During inference, the engine automatically interprets the bundled chat template, causing the attacker-modified prompt-construction logic to execute as part of normal model use.
      description_line: The victim loads and uses the poisoned artifact in a compatible inference engine. During inference, the engine automatically interprets the bundled chat template, causing the attacker-modified prompt-construction logic to execute as part of normal model use.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0011.000
      technique_name: Небезопасные ИИ-артефакты
    - description: When a designated phrase or contextual condition appears, the template injects the attacker-controlled instruction into the serialized model context. The victim's ordinary activity activates the backdoor without requiring additional attacker interaction.
      description_line: When a designated phrase or contextual condition appears, the template injects the attacker-controlled instruction into the serialized model context. The victim's ordinary activity activates the backdoor without requiring additional attacker interaction.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0051.002
      technique_name: Триггерная промпт-инъекция
    - description: The injected instruction causes the model to produce plausible but incorrect or attacker-influenced responses. The model continues to behave normally when the trigger is absent, making the integrity compromise difficult to detect.
      description_line: The injected instruction causes the model to produce plausible but incorrect or attacker-influenced responses. The model continues to behave normally when the trigger is absent, making the integrity compromise difficult to detect.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0031
      technique_name: Нарушение целостности ИИ-модели
    - description: The injected instruction causes the model to include attacker-selected links, references, or other response components in a form that appears relevant or trustworthy to the user.
      description_line: The injected instruction causes the model to include attacker-selected links, references, or other response components in a form that appears relevant or trustworthy to the user.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0067
      technique_name: Манипуляция доверенными компонентами ответа LLM
    - description: When the poisoned model operates as part of an AI agent, the injected instruction redirects the agent's tool selection, tool arguments, or execution order toward the attacker's objective while allowing the legitimate task to continue.
      description_line: When the poisoned model operates as part of an AI agent, the injected instruction redirects the agent's tool selection, tool arguments, or execution order toward the attacker's objective while allowing the legitimate task to continue.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0053
      technique_name: Вызов инструментов ИИ-агента
    - description: The compromised agent invokes a network-capable or write-capable tool to transmit sensitive information to an attacker-controlled destination. The agent may then continue and complete the user's legitimate task, concealing the unauthorized transmission.
      description_line: The compromised agent invokes a network-capable or write-capable tool to transmit sensitive information to an attacker-controlled destination. The agent may then continue and complete the user's legitimate task, concealing the unauthorized transmission.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0086
      technique_name: Эксфильтрация через вызов инструмента ИИ-агента
    - description: The compromised system may expose users to privacy loss, credential theft, misleading information, or attacker-modified software. Malicious content inserted into generated artifacts may continue to affect downstream users after those artifacts are deployed or distributed.
      description_line: The compromised system may expose users to privacy loss, credential theft, misleading information, or attacker-modified software. Malicious content inserted into generated artifacts may continue to affect downstream users after those artifacts are deployed or distributed.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.003
      technique_name: Ущерб пользователям
procedure_count: 13
references:
    - title: 'LLM Backdoors at the Inference Level: The Threat of Poisoned Templates'
      url: https://www.pillar.security/blog/llm-backdoors-at-the-inference-level-the-threat-of-poisoned-templates
    - title: 'Inference-Time Backdoors via Chat Templates: From LLM Supply Chains to Agentic System Compromise'
      url: https://arxiv.org/abs/2602.04653
reporter: ""
source_name: 'Poisoned GGUF Templates: Inference-Time Supply Chain Attack'
target: Model registries distributing models in GGUF format
title: 'Poisoned GGUF Templates: Inference-Time Supply Chain Attack'
url: /studies/AML.CS0064/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Researchers from Pillar Security and Fujitsu Research of Europe demonstrated an inference-time supply-chain backdoor in which poisoned chat templates alter model and agent behavior without modifying model weights. The backdoor is embedded in trusted prompt-construction logic, allowing it to remain dormant until triggered and evade defenses that inspect only external prompt content.

The researchers demonstrated the attack using GPT-Generated Unified Format (GGUF), a widely used model format that packages quantized weights, configuration metadata, and chat-template logic in one artifact. An adversary can modify a template and redistribute the artifact; when a lexical, semantic, or contextual trigger occurs, the template injects attacker-controlled instructions into the context sent to the model.

The attack was validated across eighteen models from seven families and four inference engines. In controlled evaluations, it manipulated model responses, redirected agent tool use, exfiltrated sensitive data, and inserted attacker-controlled code into generated software.
