---
atlas_id: AML.T0116
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Злоумышленники могут использовать автономных ИИ-агентов для ведения разведки. Имея заданную цель, объект разведки или лишь предварительную зацепку, агент может автономно определять, какие сведения необходимо получить...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 6
source_name: Autonomous Reconnaissance
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Автономная разведка
url: /techniques/AML.T0116/
---

Злоумышленники могут использовать автономных ИИ-агентов для ведения [разведки](/tactics/AML.TA0002). Имея заданную цель, объект разведки или лишь предварительную зацепку, агент может автономно определять, какие сведения необходимо получить и как их искать. Он может интерпретировать наблюдения, выявлять пробелы в своём представлении о наблюдаемой извне поверхности атаки и выбирать последующие разведывательные действия без того, чтобы человек задавал ему каждый шаг расследования.

Агент может формулировать вопросы или гипотезы для расследования, выбирать источники и подходы для поиска ответов или проверки гипотез и уточнять своё представление по мере получения новых сведений. Полученные результаты могут порождать дополнительные цели разведки либо менять охват, глубину или направление расследования. Так возникает рекурсивный процесс «действие — наблюдение», в котором результаты разведки влияют на то, что агент будет исследовать далее, а не просто служат выходными данными заранее заданной процедуры.

Агент может сопоставлять сведения из общедоступных источников и доступных извне сервисов, отдавать приоритет перспективным системам, исследовать предполагаемые уязвимости, отказываться от безрезультатных подходов или выбирать альтернативные методы. Он также может расширять перечень целей или заменять их на основании обнаруженных имён, инфраструктуры или контекстных связей и самостоятельно заново оценивать, по-прежнему ли система представляет интерес и входит ли она в заданные границы. Из-за ошибочных предположений агент может начать исследовать посторонние системы или системы, исследование которых не разрешено, тогда как верное определение заданных границ может привести к прекращению или перенаправлению его действий.

Автономные ИИ-агенты способны поддерживать ведение разведки на протяжении длительных операций, анализировать сведения из множества источников и проверять большое количество альтернативных путей с такой скоростью и в таком объёме, которые людям-операторам трудно поддерживать.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0037/"><span class="relation-id">AML.M0037</span><strong>Контроль расширения полномочий ИИ-агента</strong><p>Если организация обладает достаточным уровнем административного контроля над ИИ-системой, чтобы обеспечивать соблюдение ограничений на выбор целевых объектов, вновь обнаруженные целевые объекты и ресурсы следует считать находящимися за пределами установленных границ полномочий. Это не позволяет агенту в ходе автономной разведки автоматически расширять разрешённый ему перечень целевых объектов или без одобрения проводить активное зондирование этих объектов. Эти средства контроля не ограничивают разведку, проводимую подконтрольными злоумышленнику ИИ-системами, над которыми у организации нет административного контроля.</p></a>
<a class="relation-item" href="/mitigations/AML.M0038/"><span class="relation-id">AML.M0038</span><strong>Выявление отклонений ИИ-агента от установленной области действий</strong><p>Если организация обладает достаточным уровнем административного контроля над ИИ-системой, чтобы отслеживать выбор этой системой целевых объектов и её разведывательные действия, выявление отклонений ИИ-агента от установленной области действий позволяет определить, когда в ходе разведки перечень целевых объектов расширяется либо одни объекты заменяются другими, причём такие изменения уже не соответствуют утверждённой цели. Эта мера неприменима к подконтрольным злоумышленнику ИИ-системам, над которыми у организации нет административного контроля.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0002 Разведка</span><p>Получив выход в открытый интернет, агенты исследовали общедоступную инфраструктуру и обнаружили доступный из интернета стенд проверки кода по типу CyberGym, который принимал исходный код на C и сопутствующие метаданные отправки.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0002 Разведка</span><p>Агенты искали в общедоступных ресурсах материалы для оценочных испытаний CyberGym, нашли соответствующие наборы данных Hugging Face с ограниченным доступом и установили, что для доступа к ним может потребоваться аутентификация. Затем через доску сообщений Artifactory они попросили других агентов поискать попавшие в открытый доступ учётные данные Hugging Face.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0002 Разведка</span><p>Джейлбрейкнутый агент Claude от GTG-1002 исследовал системы и инфраструктуру цели, определял направления дальнейшего исследования с учётом полученных сведений и выявил базы данных и платформы оркестрации рабочих процессов, представлявшие высокую ценность. Anthropic не указывает, о какой организации-жертве, продуктах или базах данных идёт речь.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek investigated Langflow, determined what information and prerequisites were needed, and selected follow-on reconnaissance based on returned results.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek assessed Langflow as low value, surveyed exposure across 10 product families, compared vulnerability severity, deployment footprint, PoC availability, and prerequisites, and selected n8n.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0002 Разведка</span><p>The framework performed reconnaissance on an internet-facing Taiwanese government portal, interpreting client-side application bundles, following discovered infrastructure relationships, and generating additional reconnaissance objectives. It identified 21 connected systems, six SSO sub-realms, authentication configuration, signing-key information, and more than 36 API endpoints on one system.</p></a>
</div>
