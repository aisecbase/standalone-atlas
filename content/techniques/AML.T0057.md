---
atlas_id: AML.T0057
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленники могут создавать промпты, которые заставляют LLM раскрывать чувствительную информацию. К такой информации могут относиться приватные данные пользователей или проприетарная информация. Утекшие данные...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: LLM Data Leakage
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0010
title: Утечка данных из LLM
url: /techniques/AML.T0057/
---

Злоумышленники могут создавать промпты, которые заставляют LLM раскрывать чувствительную информацию.

К такой информации могут относиться приватные данные пользователей или проприетарная информация.

Утекшие данные могут поступать из проприетарных обучающих данных, источников данных, к которым подключена LLM, или из информации других пользователей LLM.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0010/"><span class="relation-id">AML.TA0010</span><strong>Эксфильтрация</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Тщательная оценка ИИ-модели может использоваться для выявления рисков для конфиденциальности, утечек данных и возможности раскрытия чувствительной информации.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Защитные ограничения (Guardrails) для генеративного ИИ</strong><p>Защитные ограничения могут обнаруживать чувствительные данные и персональные данные в выходных данных модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Инструкции для модели могут предписывать ей отказываться отвечать на небезопасные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Размещайте синтетические секреты или записи-канарейки в репрезентативных источниках данных и пытайтесь извлечь их через промпты, механизмы извлечения данных, инструменты и отображаемые выходные данные. Совершенствуйте разграничение полномочий, фильтрацию, изоляцию тенантов и обнаружение эксфильтрации.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0010 Эксфильтрация</span><p>Вредоносные инструкции в промпте заставляют сгенерированный ответ раскрывать чувствительные данные, например электронные письма, адреса и номера телефонов.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Раскрытие секретов через Claude Code GitHub Action</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0010 Эксфильтрация</span><p>Claude emitted the Anthropic API key after removing its `sk-ant-` prefix. The transformation prevented GitHub&#39;s secret scanner from recognizing the credential, while the researchers could reconstruct the original key by restoring the prefix.</p></a>
</div>
