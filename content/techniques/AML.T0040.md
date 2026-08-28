---
atlas_id: AML.T0040
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут получить доступ к модели через легитимный доступ к API инференса. Доступ к API инференса может быть для злоумышленника источником информации (выявление онтологии ИИ-модели, выявление семейства...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 8
source_name: AI Model Inference API Access
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0000
title: Доступ к API инференса ИИ-модели
url: /techniques/AML.T0040/
---

Злоумышленники могут получить доступ к модели через легитимный доступ к API инференса. Доступ к API инференса может быть для злоумышленника источником информации ([выявление онтологии ИИ-модели](/techniques/AML.T0013), [выявление семейства ИИ-модели](/techniques/AML.T0014)), средством подготовки атаки ([проверка атаки](/techniques/AML.T0042), [создание состязательных данных](/techniques/AML.T0043)) или способом передать данные в целевую систему для воздействия ([уклонение от ИИ-модели](/techniques/AML.T0015), [нарушение целостности ИИ-модели](/techniques/AML.T0031)).

Многие системы полагаются на одни и те же модели, доступные через API инференса, поэтому они наследуют общие уязвимости. Это особенно характерно для фундаментальных моделей, обучение которых требует запретительно больших ресурсов. Злоумышленники могут использовать доступ к API моделей, чтобы выявлять уязвимости, например джейлбрейки или галлюцинации, а затем атаковать приложения, использующие те же модели.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0000/"><span class="relation-id">AML.TA0000</span><strong>Доступ к ИИ-модели</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Злоумышленники могут использовать неограниченный доступ к API, чтобы получить сведения о продакшен-системе, подготовить атаки и внедрить в систему вредоносные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помогать аудировать использование API модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи использовали публично доступное приложение не по назначению: отправляли запросы к модели и получали машинно переведенные пары предложений в качестве обучающих данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Команда использовала открытый API для доступа к целевой модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0011/"><span class="relation-id">AML.CS0011</span><strong>Обход ИИ на периферии Microsoft</strong><span class="relation-meta">Актор: Azure Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Используя общедоступную версию ML-модели, команда начала отправлять запросы и анализировать ответы, то есть результаты инференса модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Команда получила доступ к API инференса целевой модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>На протяжении упражнения исследователи использовали публичный API ChatGPT.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Исследователи используют доступ к публичному API GenAI-модели, на которой работает целевая почтовая система с RAG.</p></a>
<a class="relation-item" href="/studies/AML.CS0056/"><span class="relation-id">AML.CS0056</span><strong>Кампании по дистилляции моделей, нацеленные на Anthropic Claude</strong><span class="relation-meta">Актор: DeepSeek, Moonshot AI, MiniMax / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>ИИ-лаборатории обращались к API инференса Claude через примерно 24 000 поддельных учетных записей в совокупности.</p></a>
<a class="relation-item" href="/studies/AML.CS0057/"><span class="relation-id">AML.CS0057</span><strong>Storm-2139: обход гардрейлов Azure OpenAI</strong><span class="relation-meta">Актор: Storm-2139 / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Похищенные учетные данные давали доступ к Azure OpenAI Service, позволяя злоумышленникам и их клиентам отправлять промпты и генерировать контент. Инструмент Storm-2139 de3u использовался как фронтенд для этого доступа.</p></a>
</div>
