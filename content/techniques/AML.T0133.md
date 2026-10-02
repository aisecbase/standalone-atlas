---
atlas_id: AML.T0133
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Злоумышленники могут взаимодействовать с ИИ-агентом во время его работы, чтобы выявить доступные ему возможности. Для этого им не требуется доступ к конфигурации агента. Прямое взаимодействие с агентом может раскрыть...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Agentic AI
procedure_count: 0
source_name: Discover AI Agent Runtime Capabilities
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление возможностей ИИ-агента во время работы
url: /techniques/AML.T0133/
---

Злоумышленники могут взаимодействовать с ИИ-агентом во время его работы, чтобы выявить доступные ему возможности. Для этого им не требуется доступ к конфигурации агента. Прямое взаимодействие с агентом может раскрыть зарегистрированные у него инструменты и принимаемые ими параметры, действия, которые он может выполнять, доступные ему ресурсы, а также то, от чьего имени и с какими правами он действует. О возможностях агента также можно судить косвенно: отправлять различные запросы и наблюдать, какие выполнены успешно, какие завершились ошибкой, а какие агент отказался выполнять.

ИИ-агенты часто подключены к корпоративным ресурсам, инструментам и базам данных либо встроены в SaaS-платформы и для выполнения своих функций имеют разрешения действовать от имени пользователей. Выявив работающего агента, к которому у них есть доступ, злоумышленники могут исследовать его поверхность атак: проверять его функциональность, выявлять его инструменты и возможности, доступные ему знания, а также встроенные учётные данные и разрешения.

Такое исследование часто раскрывает полный набор инструментов ИИ-агента и сведения о его конфигурации, открывая дополнительные возможности эксплуатации с помощью техники [Вызов инструментов ИИ-агента](/techniques/AML.T0053). Полученные сведения облегчают дальнейшую эксплуатацию, включая [Первичный доступ](/tactics/AML.TA0004), [Закрепление](/tactics/AML.TA0006), [Повышение привилегий](/tactics/AML.TA0012) и [Эксфильтрацию](/tactics/AML.TA0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>Interactive agentic decoys can present a plausible but fake toolset and capabilities. When an adversary enumerates the agent&#39;s capabilities at runtime, the honeypot can capture which tools and privileges they probe for and how they attempt to exploit it, while planted decoy credentials can also be used as canary tokens.</p></a>
</div>
