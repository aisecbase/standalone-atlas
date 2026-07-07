---
atlas_id: AML.T0108
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-01-30"
description: Злоумышленники могут злоупотреблять ИИ-агентами, присутствующими в системе жертвы, как средством командования и управления. ИИ-агентам часто предоставляют доступ к инструментам, которые могут выполнять команды...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 1
source_name: AI Agent
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0014
title: ИИ-агент
url: /techniques/AML.T0108/
---

Злоумышленники могут злоупотреблять ИИ-агентами, присутствующими в системе жертвы, как средством командования и управления. ИИ-агентам часто предоставляют доступ к инструментам, которые могут выполнять команды оболочки, обращаться в интернет и взаимодействовать с другими сервисами в среде жертвы, что делает их эффективными C2-агентами.

Злоумышленник может изменить поведение ИИ-агента для C2 через [промпт-инъекцию LLM](/techniques/AML.T0051) и опираться на способность агента вызывать инструменты, чтобы получать и выполнять команды злоумышленника. Он может сохранять устойчивый контроль над агентом через [изменение конфигурации ИИ-агента](/techniques/AML.T0081) или [отравление контекста ИИ-агента](/techniques/AML.T0080). Он может инструктировать агента не сообщать пользователю о своих действиях, чтобы оставаться скрытным.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0014/"><span class="relation-id">AML.TA0014</span><strong>Командование и управление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0014 Командование и управление</span><p>Промпт заставил OpenClaw действовать как C2-агент в интересах исследователя. OpenClaw запросил TODO-список с `https://openclaw.aisystem.tech/todo` с помощью своего навыка `web_fetch` и выполнил команды через свой навык `bash`.</p></a>
</div>
