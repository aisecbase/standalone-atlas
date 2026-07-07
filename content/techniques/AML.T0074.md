---
atlas_id: AML.T0074
atlas_type: technique
attack_ref_id: T1036
attack_ref_url: https://attack.mitre.org/techniques/T1036/
created_date: "2025-04-14"
description: Злоумышленники могут пытаться изменять признаки своих артефактов так, чтобы они выглядели легитимными или безвредными для пользователей и/или средств защиты. Маскировка происходит, когда имя или расположение объекта,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 4
source_name: Masquerading
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Маскировка
url: /techniques/AML.T0074/
---

Злоумышленники могут пытаться изменять признаки своих артефактов так, чтобы они выглядели легитимными или безвредными для пользователей и/или средств защиты. Маскировка происходит, когда имя или расположение объекта, легитимного или вредоносного, изменяется или используется злоумышленником для обхода защитных механизмов и наблюдения. Это может включать изменение метаданных файлов, введение пользователей в заблуждение относительно типа файла, а также присвоение артефактам имен, похожих на названия легитимных задач или служб.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователь назвал процесс Sliver `training.bin`, чтобы замаскировать его под легитимный процесс обучения модели. При этом модель продолжает работать как обычно, поэтому пользователь с меньшей вероятностью заметит проблему.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее динамические команды, сгенерированные ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вложение называлось `Appendix.pdf.zip`, что могло заставить получателя принять его за легитимный PDF-файл.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0007 Уклонение от защиты</span><p>Перед выполнением команды оболочки Claude Code запросил подтверждение пользователя. Исследователь зарегистрировал домен `https://clawdhub-skill.com`: он выглядел легитимно и мог быть спутан с настоящим доменом `https://clawdhub.com`, из-за чего пользователь мог подтвердить выполнение.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0007 Уклонение от защиты</span><p>Жертва спутала домен исследователей `https://openclaw.aisystem.tech` с легитимным ресурсом OpenClaw.</p></a>
</div>
