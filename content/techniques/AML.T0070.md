---
atlas_id: AML.T0070
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут внедрять вредоносное содержимое в данные, индексируемые системой RAG (генерации, дополненной извлечением), чтобы вредоносное содержимое попадало в будущий диалог через результаты поиска на основе...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: RAG Poisoning
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0006
title: Отравление RAG
url: /techniques/AML.T0070/
---

Злоумышленники могут внедрять вредоносное содержимое в данные, индексируемые системой RAG (генерации, дополненной извлечением), чтобы вредоносное содержимое попадало в будущий диалог через результаты поиска на основе RAG. Для этого они могут разместить манипулированные документы в месте, которое индексирует RAG (см. [Сбор целей, индексируемых RAG](/techniques/AML.T0064)).

Содержимое может быть нацелено так, чтобы всегда появляться в результатах поиска по конкретному пользовательскому запросу. Содержимое злоумышленника может включать ложную или вводящую в заблуждение информацию. Оно также может включать промпт-инъекции с вредоносными инструкциями или ложные записи RAG.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Защитные ограничения (Guardrails) для генеративного ИИ</strong><p>Apply retrieval guardrails to reject untrusted, malicious, irrelevant, or unsupported RAG content.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Размещайте контролируемое вредоносное или вводящее в заблуждение содержимое в репрезентативных источниках поступления данных. Совершенствуйте механизмы допуска источников, отслеживание происхождения данных, валидацию содержимого, средства контроля индексирования и фильтрацию при извлечении данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0006 Закрепление</span><p>Исследователи добились закрепления в системе жертвы, поскольку вредоносный промпт выполнялся каждый раз, когда извлекалась отравленная RAG-запись с поддельными банковскими реквизитами.</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0006 Закрепление</span><p>Исследователь создает публичный канал Slack и отправляет в него вредоносное содержимое: текст для извлечения и промпт. Поскольку Slack AI индексирует сообщения в публичных каналах, вредоносное сообщение добавляется в его RAG-базу данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0006 Закрепление</span><p>Письмо было автоматически загружено в базу данных RAG, доступную пайплайну извлечения Copilot.</p></a>
</div>
