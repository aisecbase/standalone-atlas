---
actor: Pillar Security, Fujitsu Research of Europe
atlas_id: AML.CS0064
atlas_type: case-study
case_study_type: exercise
description: 'Исследователи из Pillar Security и Fujitsu Research of Europe продемонстрировали бэкдор в цепочке поставок, срабатывающий во время инференса: отравленные шаблоны чата изменяют поведение модели и агента без изменения...'
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
      technique_name: Модели
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
title: 'Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса'
url: /studies/AML.CS0064/
---

Исследователи из Pillar Security и Fujitsu Research of Europe продемонстрировали бэкдор в цепочке поставок, срабатывающий во время инференса: отравленные шаблоны чата изменяют поведение модели и агента без изменения весов модели. Бэкдор встроен в доверенную логику формирования промпта, что позволяет ему оставаться неактивным до срабатывания триггера и обходить средства защиты, анализирующие только внешнее содержимое промпта.

Исследователи продемонстрировали атаку с использованием GPT-Generated Unified Format (GGUF) — широко используемого формата моделей, объединяющего в одном артефакте квантованные веса, метаданные конфигурации и логику шаблона чата. Злоумышленник может изменить шаблон и повторно распространить артефакт; при срабатывании лексического, семантического или контекстного триггера шаблон внедряет подконтрольные злоумышленнику инструкции в контекст, передаваемый модели.

Атака была проверена на восемнадцати моделях из семи семейств с использованием четырёх движков инференса. В ходе контролируемых испытаний она приводила к манипулированию ответами модели, перенаправлению вызовов инструментов агентом, эксфильтрации чувствительных данных и внедрению подконтрольного злоумышленнику кода в генерируемое ПО.
