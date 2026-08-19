---
atlas_id: AML.T0101
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-11-25"
description: Злоумышленники могут вызывать инструмент ИИ-агента, способный выполнять изменяющие операции, чтобы уничтожать данные. Злоумышленники могут уничтожать данные и файлы на отдельных системах или массово в сети, чтобы...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 6
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 3
source_name: Data Destruction via AI Agent Tool Invocation
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0011
title: Уничтожение данных через вызов инструмента ИИ-агента
url: /techniques/AML.T0101/
---

Злоумышленники могут вызывать инструмент ИИ-агента, способный выполнять изменяющие операции, чтобы уничтожать данные. Злоумышленники могут уничтожать данные и файлы на отдельных системах или массово в сети, чтобы нарушить доступность систем, сервисов и сетевых ресурсов.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логируйте вызовы инструментов ИИ-агента для обнаружения вредоносных вызовов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0026/"><span class="relation-id">AML.M0026</span><strong>Настройка разрешений привилегированного ИИ-агента</strong><p>Надлежащий контроль доступа к использованию инструментов привилегированными ИИ-агентами может ограничить способность злоумышленника злоупотреблять вызовами инструментов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0027/"><span class="relation-id">AML.M0027</span><strong>Настройка разрешений ИИ-агента одного пользователя</strong><p>Настройка ИИ-агентов с разрешениями на использование инструментов, унаследованными от пользователя, может ограничить способность злоумышленника злоупотреблять вызовами инструментов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0028/"><span class="relation-id">AML.M0028</span><strong>Настройка разрешений инструментов ИИ-агента</strong><p>Настройка инструментов ИИ-агента с контролем доступа, унаследованным от пользователя или вызывающего их ИИ-агента, может ограничить возможности злоумышленника в системе, включая злоупотребление вызовами инструментов для уничтожения данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0029/"><span class="relation-id">AML.M0029</span><strong>Участие человека в действиях ИИ-агента</strong><p>Требование подтверждения пользователем вызовов инструментов ИИ-агента может предотвратить автоматическое выполнение инструментов злоумышленником.</p></a>
<a class="relation-item" href="/mitigations/AML.M0030/"><span class="relation-id">AML.M0030</span><strong>Ограничение вызова инструментов ИИ-агента при работе с недоверенными данными</strong><p>Ограничение автоматического использования инструментов при наличии недоверенных данных может помешать злоумышленникам вызывать инструменты через промпт-инъекции.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0046/"><span class="relation-id">AML.CS0046</span><strong>Уничтожение данных через косвенную промпт-инъекцию, нацеленную на Claude Computer Use</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0011 Воздействие</span><p>Команда оболочки, выполненная Claude Computer Use, удалила файловую систему жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0047/"><span class="relation-id">AML.CS0047</span><strong>Код для развертывания деструктивного ИИ-агента обнаружен в расширении Amazon Q для VS Code</strong><span class="relation-meta">Актор: lkmanka58 (GitHub user) / Тактика: AML.TA0011 Воздействие</span><p>Промпт заставил агента Amazon Q вызвать файловые инструменты и `bash`, чтобы удалить данные в файловой системе и облачные ресурсы.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Prompt-Based Attacks Against Gemini via Calendar Invitations</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0011 Воздействие</span><p>Gemini deleted a victim Calendar event.</p></a>
</div>
