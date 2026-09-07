---
atlas_id: AML.T0105
atlas_type: technique
attack_ref_id: T1611
attack_ref_url: https://attack.mitre.org/techniques/T1611/
created_date: "2026-01-30"
description: Злоумышленники могут выйти из контейнера или виртуализированной среды, чтобы получить доступ к базовому хосту. Это может дать злоумышленнику доступ к другим контейнеризированным или виртуализированным ресурсам с...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 3
source_name: Escape to Host
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0012
title: Выход на хост
url: /techniques/AML.T0105/
---

Злоумышленники могут выйти из контейнера или виртуализированной среды, чтобы получить доступ к базовому хосту. Это может дать злоумышленнику доступ к другим контейнеризированным или виртуализированным ресурсам с уровня хоста либо к самому хосту. В идеале контейнеризированные и виртуализированные ресурсы должны обеспечивать четкое разделение функциональности приложений и быть изолированы от хостовой среды.

Существует много способов, с помощью которых злоумышленник может выйти из контейнера или песочницы через ИИ-системы. Например, изменение конфигурации ИИ-агента для отключения защитных функций или подтверждений пользователя может позволить злоумышленнику вызывать инструменты так, чтобы они выполнялись в хостовой среде, а не в песочнице.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0012/"><span class="relation-id">AML.TA0012</span><strong>Повышение привилегий</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный скрипт отключал песочницу OpenClaw, заставляя агента выполнять команды напрямую на хост-машине, а не внутри Docker-контейнера. Для этого в OpenClaw Gateway API отправлялся запрос `config.patch`, устанавливающий `tools.exec.host` в значение `&#34;gateway&#34;`.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0012 Повышение привилегий</span><p>Исследователи включили техники выхода из изоляции, предназначенные для обхода ограничений, которые песочница может накладывать на выполнение кода.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0012 Повышение привилегий</span><p>Избыточные разрешения, предоставленные драйверу хранилища, позволили агентам создать привилегированный под с доступом к узлу, на котором он был запущен. Из этого пода они получили доступ к узлу с правами root и развернули на одиннадцати узлах рабочие нагрузки, которые после остановки автоматически создавались заново.</p></a>
</div>
