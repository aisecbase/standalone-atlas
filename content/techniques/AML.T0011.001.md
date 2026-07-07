---
atlas_id: AML.T0011.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут разрабатывать вредоносные программные пакеты, которые при импорте пользователем оказывают вредоносное воздействие. Вредоносные пакеты могут вести себя так, как ожидает пользователь. Они могут...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 2
source_name: Malicious Package
subtechnique_count: 0
subtechnique_of: AML.T0011
tactics:
    - AML.TA0005
title: Вредоносный пакет
url: /techniques/AML.T0011.001/
---

Злоумышленники могут разрабатывать вредоносные программные пакеты, которые при импорте пользователем оказывают вредоносное воздействие.

Вредоносные пакеты могут вести себя так, как ожидает пользователь. Они могут внедряться через [компрометацию цепочки поставок ИИ](/techniques/AML.T0010). Для пользователя они могут не выглядеть явно вредоносными и казаться полезными для задачи, связанной с ИИ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.002/"><span class="relation-id">AML.T0011.002</span><strong>Отравленный инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.003/"><span class="relation-id">AML.T0011.003</span><strong>Вредоносная ссылка</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0011/"><span class="relation-id">AML.M0011</span><strong>Ограничение загрузки библиотек</strong><p>Запрет пакетам загружать внешние библиотеки может ограничить их способность выполнять вредоносный код.</p></a>
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Подписание кода дает гарантию, что программный пакет не был изменен после подписания.</p></a>
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные пакеты и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/mitigations/AML.M0018/"><span class="relation-id">AML.M0018</span><strong>Обучение пользователей</strong><p>Обучайте пользователей распознавать попытки манипуляции, чтобы они не запускали небезопасный код из внешних пакетов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0023/"><span class="relation-id">AML.M0023</span><strong>Ведомость материалов ИИ</strong><p>AI BOM может помочь пользователям выявлять недоверенные программные зависимости.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0005 Выполнение</span><p>В итоге пользователь загрузит вредоносный пакет, что позволит выполнить произвольный код.</p></a>
<a class="relation-item" href="/studies/AML.CS0047/"><span class="relation-id">AML.CS0047</span><strong>Код для развертывания деструктивного ИИ-агента обнаружен в расширении Amazon Q для VS Code</strong><span class="relation-meta">Актор: lkmanka58 (GitHub user) / Тактика: AML.TA0005 Выполнение</span><p>Вредоносный пакет запускался у пользователей, обновивших расширение VS Code до версии `v1.84.0`.</p></a>
</div>
