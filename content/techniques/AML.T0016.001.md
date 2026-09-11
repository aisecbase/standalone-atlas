---
atlas_id: AML.T0016.001
atlas_type: technique
attack_ref_id: T1588.002
attack_ref_url: https://attack.mitre.org/techniques/T1588/002/
created_date: "2021-05-13"
description: Злоумышленники могут искать и получать программные инструменты для поддержки своих операций. ПО, разработанное для легитимного использования, может быть переиспользовано злоумышленником во вредоносных целях....
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 6
source_name: Software Tools
subtechnique_count: 0
subtechnique_of: AML.T0016
tactics:
    - AML.TA0003
title: Программные инструменты
url: /techniques/AML.T0016.001/
---

Злоумышленники могут искать и получать программные инструменты для поддержки своих операций.

ПО, разработанное для легитимного использования, может быть переиспользовано злоумышленником во вредоносных целях.

Злоумышленник может изменять или адаптировать программные инструменты для достижения своих целей.

Программные инструменты, используемые для поддержки атак на ИИ-системы, не обязательно сами основаны на ИИ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016/"><span class="relation-id">AML.T0016</span><strong>Получение средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.003/"><span class="relation-id">AML.T0016.003</span><strong>Эксплойты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.004/"><span class="relation-id">AML.T0016.004</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники получили кастомизированные Android ROM и приложение виртуальной камеры.</p></a>
<a class="relation-item" href="/studies/AML.CS0030/"><span class="relation-id">AML.CS0030</span><strong>LLM-джекинг</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники получили [keychecker](https://github.com/cunnymessiah/keychecker) — инструмент массовой проверки ключей для различных ИИ-сервисов, который может проверять действительность ключа и получать отдельные атрибуты учетной записи, например баланс аккаунта и доступные модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0030/"><span class="relation-id">AML.CS0030</span><strong>LLM-джекинг</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Затем злоумышленники использовали [OAI Reverse Proxy](https://gitgud.io/khanon/oai-reverse-proxy), чтобы развернуть обратный прокси-сервис для украденных LLM-ресурсов. Этот прокси-сервис можно было использовать для продажи доступа киберпреступникам, которые могли эксплуатировать LLM в вредоносных целях.</p></a>
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи получили [Open Broadcaster Software (OBS)](https://obsproject.com), которое может транслировать видеопоток по сети.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0003 Подготовка ресурсов</span><p>GTG-1002 получила сетевые сканеры, фреймворки для эксплуатации баз данных, средства взлома паролей, утилиты для анализа бинарных файлов и другие инструменты, которые стали доступны её джейлбрейкнутому агенту Claude.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Злоумышленник использовал Hermes Agent на базе DeepSeek при попытках эксплуатации Langflow и n8n</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник получил и настроил Hermes Agent как фреймворк для проведения атак вместе со скриптами и обычными утилитами для сканирования и эксплуатации уязвимостей. Hermes обеспечивал доступ к терминалу, управление со стороны оператора через Telegram и систему навыков.</p></a>
</div>
