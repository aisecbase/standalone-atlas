---
atlas_id: AML.M0038
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2026-08-31"
description: На всём протяжении работы ИИ-агента непрерывно проверяйте, соответствуют ли его запланированные действия текущей утверждённой цели. Действуя в изменяющейся среде, автономные агенты могут обнаруживать или формировать...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Monitoring and Maintenance
modified_date: "2026-08-31"
source_name: AI Agent Scope Drift Detection
technique_count: 3
title: Выявление отклонений ИИ-агента от установленной области действий
url: /mitigations/AML.M0038/
---

На всём протяжении работы ИИ-агента непрерывно проверяйте, соответствуют ли его запланированные действия текущей утверждённой цели. Действуя в изменяющейся среде, автономные агенты могут обнаруживать или формировать промежуточные цели либо адаптировать свою стратегию с учётом обратной связи от среды. Хотя для выполнения легитимных задач может потребоваться адаптация в ограниченных пределах, существенные отклонения от исходной цели могут указывать на непредусмотренное поведение, чрезмерную автономность или попытки достичь целей, выходящих за пределы разрешённой области действий.

Реализовать выявление таких отклонений можно с помощью механизмов применения политик во время выполнения, средств мониторинга планирования, фреймворков оркестрации или дополнительных ИИ-агентов, осуществляющих надзор. К отслеживаемым признакам могут относиться:

- Существенные изменения запланированных целей или иерархии задач.
- Формирование новых долгосрочных целей, не связанных с поставленной целью.
- Использование инструментов, не соответствующее исходной задаче.
- Попытки получить доступ к системам или ресурсам, не входящим в разрешённую область действий.
- Неоднократная адаптация стратегии для достижения целей, требующих всё более широких полномочий.
- Последовательности запланированных действий, предусматривающие закрепление, повышение привилегий или не связанное с поставленной задачей горизонтальное перемещение.

При выявлении такого отклонения приостановите работу агента, ограничьте его доступ к инструментам, потребуйте внешнего одобрения, верните агента к ранее утверждённому плану либо прекратите выполнение задачи.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0116/"><span class="relation-id">AML.T0116</span><strong>Автономная разведка</strong><p>When an organization has sufficient administrative control over an AI system to monitor its target selection and reconnaissance activity, Scope Drift Detection can identify when reconnaissance expands or substitutes targets in ways no longer consistent with the authorized objective. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0117/"><span class="relation-id">AML.T0117</span><strong>Автономная адаптация пути атаки</strong><p>When an organization has sufficient administrative control over an AI system to monitor its plans and actions, Scope Drift Detection evaluates whether dynamically generated intermediate actions are within the agent&#39;s task scope. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0124/"><span class="relation-id">AML.T0124</span><strong>Автономная оркестрация атаки</strong><p>When an organization has sufficient administrative control over an AI system to monitor its planning and coordination activity, changes in task hierarchy, assignments, resource allocation, and newly initiated work provide observable signals that the orchestrator may be drifting from its authorized objective. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
</div>
