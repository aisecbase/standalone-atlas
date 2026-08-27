---
atlas_id: AML.T0049
atlas_type: technique
attack_ref_id: T1190
attack_ref_url: https://attack.mitre.org/techniques/T1190/
created_date: "2023-02-28"
description: Злоумышленники могут пытаться использовать слабое место в доступном из интернета компьютере или программе с помощью ПО, данных или команд, чтобы вызвать непреднамеренное или непредусмотренное поведение. Слабое место в...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 5
source_name: Exploit Public-Facing Application
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
title: Эксплуатация приложения, доступного из интернета
url: /techniques/AML.T0049/
---

Злоумышленники могут пытаться использовать слабое место в доступном из интернета компьютере или программе с помощью ПО, данных или команд, чтобы вызвать непреднамеренное или непредусмотренное поведение. Слабое место в системе может быть ошибкой, сбоем или проектной уязвимостью. Такие приложения часто являются веб-сайтами, но могут также включать базы данных (например, SQL), стандартные сервисы (например, SMB или SSH), протоколы администрирования и управления сетевыми устройствами (например, SNMP и Smart Install), а также любые другие приложения с открытыми сокетами, доступными из интернета, например веб-серверы и связанные сервисы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay: захват кластеров Ray, доступных из интернета</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>После обнаружения открытых кластеров Ray злоумышленники могут использовать Jobs API для запуска заданий на доступных кластерах. Jobs API не поддерживает авторизацию, поэтому любой, у кого есть сетевой доступ к кластеру, может удаленно выполнить произвольный код.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи смогли воспользоваться ошибочно настроенными реестрами и скачать контейнерные образы без аутентификации. В общей сложности они скачали несколько терабайт данных, содержащих более 20 000 образов.</p></a>
<a class="relation-item" href="/studies/AML.CS0030/"><span class="relation-id">AML.CS0030</span><strong>LLM-джекинг</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Злоумышленники воспользовались уязвимой версией Laravel ([CVE-2021-3129](https://www.cve.org/CVERecord?id=CVE-2021-3129)), чтобы получить первичный доступ к системам жертв.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учетные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователь воспользовался ошибочной конфигурацией прокси на сервере управления ClawdBot и получил доступ к интерфейсам управления с включенной аутентификацией.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи нацеливались на публично доступные приложения, где ИИ-агент принимает пользовательский ввод, чтобы через него выполнить свои промпты.</p></a>
</div>
