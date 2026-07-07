---
atlas_id: AML.T0062
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут отправлять запросы большим языковым моделям и выявлять галлюцинированные сущности. Они могут запрашивать программные пакеты, команды, URL-адреса, названия организаций или адреса электронной почты...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Discover LLM Hallucinations
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление галлюцинированных сущностей LLM
url: /techniques/AML.T0062/
---

Злоумышленники могут отправлять запросы большим языковым моделям и выявлять галлюцинированные сущности.

Они могут запрашивать программные пакеты, команды, URL-адреса, названия организаций или адреса электронной почты и находить галлюцинации, за которыми не стоит реальный источник. Обнаруженные галлюцинированные сущности дают злоумышленнику потенциальные цели для [публикации галлюцинированных сущностей](/techniques/AML.T0060). Было показано, что разные LLM могут выдавать одни и те же галлюцинации, поэтому сущности, найденные и эксплуатируемые злоумышленником, могут затрагивать пользователей других LLM.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение количества запросов к ИИ-модели</strong><p>Ограничение количества запросов к модели ограничивает или замедляет способность злоумышленника выявлять возможные галлюцинации.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Защитные ограничения (Guardrails) для генеративного ИИ</strong><p>Защитные ограничения могут помогать блокировать галлюцинированный контент в выходных данных модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Инструкции могут указывать модели избегать генерации галлюцинированного контента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивание модели может помогать уводить модель от галлюцинированного контента.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0008 Выявление</span><p>Исследователи просили ChatGPT предложить программные пакеты и выявляли среди рекомендаций галлюцинации — пакеты, которых нет в публичном репозитории. Например, на вопрос &#34;how to upload a model to huggingface?&#34; модель предложила установить пакет `huggingface-cli` командой `pip install huggingface-cli`. Такого пакета в PyPI не существовало; реальный CLI-инструмент Hugging Face входит в пакет `huggingface_hub`.</p></a>
</div>
