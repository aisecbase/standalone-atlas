---
atlas_id: AML.T0074
atlas_type: technique
attack_ref_id: T1036
attack_ref_url: https://attack.mitre.org/techniques/T1036/
created_date: "2025-04-14"
description: Злоумышленники могут пытаться изменять признаки своих артефактов, чтобы они выглядели легитимными или безвредными для пользователей и/или средств защиты. Маскировка имеет место, когда имя или расположение легитимного...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-07-31"
platforms:
    - Enterprise
procedure_count: 6
source_name: Masquerading
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Маскировка
url: /techniques/AML.T0074/
---

Злоумышленники могут пытаться изменять признаки своих артефактов, чтобы они выглядели легитимными или безвредными для пользователей и/или средств защиты. Маскировка имеет место, когда имя или расположение легитимного либо вредоносного объекта изменяют или используют для уклонения от защиты и наблюдения. Это может включать изменение метаданных файлов, введение пользователей в заблуждение относительно типа файла и присвоение объектам имён легитимных задач или служб.

В публичных реестрах артефактов, таких как Hugging Face, злоумышленники могут повторно зарегистрировать или воссоздать ранее доверенное пространство имён либо путь после удаления или переименования соответствующего идентификатора. После этого при разрешении устаревших идентификаторов без зафиксированной версии или ревизии системы могут получить подконтрольные злоумышленнику артефакты, которые выглядят как тот самый артефакт, которому изначально доверяли.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0011/"><span class="relation-id">AML.M0011</span><strong>Ограничение загрузки библиотек</strong><p>Restrict library loading to trusted locations and approved libraries so disguised malicious libraries cannot be loaded as legitimate dependencies.</p></a>
<a class="relation-item" href="/mitigations/AML.M0025/"><span class="relation-id">AML.M0025</span><strong>Поддержание происхождения наборов данных ИИ</strong><p>Record dataset sources and modification history so datasets falsely presented as trusted can be identified.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователь назвал процесс Sliver `training.bin`, чтобы замаскировать его под легитимный процесс обучения модели. При этом модель продолжает работать как обычно, поэтому пользователь с меньшей вероятностью заметит проблему.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вложение называлось `Appendix.pdf.zip`, что могло заставить получателя принять его за легитимный PDF-файл.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0007 Уклонение от защиты</span><p>Перед выполнением команды оболочки Claude Code запросил подтверждение пользователя. Исследователь зарегистрировал домен `https://clawdhub-skill.com`: он выглядел легитимно и мог быть спутан с настоящим доменом `https://clawdhub.com`, из-за чего пользователь мог подтвердить выполнение.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0007 Уклонение от защиты</span><p>Жертва спутала домен исследователей `https://openclaw.aisystem.tech` с легитимным ресурсом OpenClaw.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0007 Уклонение от защиты</span><p>Злоумышленник делает так, чтобы изменённый артефакт выглядел эквивалентным легитимной модели. В отсутствие триггера артефакт сохраняет ожидаемое поведение, а вредоносная логика скрыта среди легитимных элементов форматирования шаблона и легитимной управляющей логики.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Атака на цепочку поставок через повторное использование пространства имён модели</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0007 Уклонение от защиты</span><p>Воссоздав пространство имён, специалисты Unit 42 добились того, чтобы вредоносный артефакт выглядел как ранее доверенная модель. В случае моделей с переданным владением повторная регистрация старого пространства имён вытесняла прежнее перенаправление к новому расположению легитимной модели.</p></a>
</div>
