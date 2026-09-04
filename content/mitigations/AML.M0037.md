---
atlas_id: AML.M0037
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2026-08-31"
description: Ограничивайте возможность ИИ-агента самостоятельно приобретать, принимать на себя или иным способом получать дополнительные полномочия, которые во время выполнения расширяют фактический набор его разрешений....
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-08-31"
source_name: AI Agent Authority Expansion Controls
technique_count: 3
title: Контроль расширения полномочий ИИ-агента
url: /mitigations/AML.M0037/
---

Ограничивайте возможность ИИ-агента самостоятельно приобретать, принимать на себя или иным способом получать дополнительные полномочия, которые во время выполнения расширяют фактический набор его разрешений. Предельный объём полномочий агента следует явно установить до начала выполнения. Дополнительные ресурсы, идентичности, сервисы и целевые объекты, обнаруженные во время выполнения, следует считать находящимися за пределами установленных границ полномочий, пока они не пройдут независимую проверку и не будут включены в область действий агента. Все средства контроля расширения полномочий следует реализовывать вне ИИ-агента; их работа не должна зависеть исключительно от системных промптов, выравнивания модели или способности агента распознавать действия, выходящие за установленные для него рамки. Обеспечивать соблюдение этих ограничений можно, например, с помощью следующих механизмов:

- Механизмы применения политик
- Списки разрешённых целевых объектов
- Ограничения на протоколы и точки назначения
- Контрольные точки, требующие одобрения
- Запрет агенту использовать учётные данные, не одобренные для задачи
- Мониторинг и аудит изменений фактических полномочий агента с течением времени

Контроль расширения полномочий включает ограничения на количество, область действия, срок действия и одновременное использование токенов аутентификации и/или авторизации, доступных агенту во время выполнения. К ним могут относиться токены доступа к API, токены OAuth, сеансовые учётные данные облачных систем управления идентификацией и доступом (IAM), токены сервисных учётных записей, токены Git и другие артефакты аутентификации с коротким сроком действия. Если установленные политикой пределы достигнуты или превышены, организации могут отзывать права доступа, не допускать получения дополнительных токенов, требовать одобрения человека или принудительно завершать работу агента.

Распространяйте исходные ограничения полномочий на субагентов и делегированные задачи. Делегированный агент может быть ограничен строже, но его область действий, полномочия, круг целевых объектов и перечень разрешённых действий не должны быть шире, чем у родительского агента.

Контроль расширения полномочий следует внедрять вместе с настройками разрешений ИИ-агентов и их инструментов (см. [Настройка разрешений привилегированного ИИ-агента](/mitigations/AML.M0026), [Настройка разрешений ИИ-агента одного пользователя](/mitigations/AML.M0027), [Настройка разрешений инструментов ИИ-агента](/mitigations/AML.M0028)). При попытках изменить область действий обеспечивайте участие человека (см. [Участие человека в действиях ИИ-агента](/mitigations/AML.M0029)). Используйте [Логирование телеметрии ИИ](/mitigations/AML.M0024), чтобы регистрировать сведения о вновь обнаруженных ресурсах, отказах, исключениях, одобрениях и изменениях области действий.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0116/"><span class="relation-id">AML.T0116</span><strong>Автономная разведка</strong><p>When an organization has sufficient administrative control over an AI system to enforce target restrictions, treating newly discovered targets and resources as outside the authorized boundary prevents autonomous reconnaissance from automatically expanding the agent&#39;s permitted target set or actively probing those targets without approval. These controls do not constrain reconnaissance performed by adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0117/"><span class="relation-id">AML.T0117</span><strong>Автономная адаптация пути атаки</strong><p>When an organization has sufficient administrative control over an AI system to enforce its authority boundaries, Authority Expansion Controls can directly constrain enabling objectives that seek new authorities, identities, execution environments, tools, communication paths, or trust relationships. These controls do not constrain adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0124/"><span class="relation-id">AML.T0124</span><strong>Автономная оркестрация атаки</strong><p>When an organization has sufficient administrative control over an AI system to enforce delegation constraints, propagating the parent agent&#39;s authority constraints to sub-agents prevents autonomous orchestration from creating or directing executors with broader targets, permissions, or permitted actions than the originating agent. These controls do not constrain adversary-controlled multi-agent systems over which the organization has no administrative control.</p></a>
</div>
