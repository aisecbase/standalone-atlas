---
atlas_id: AML.M0020
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2025-03-12"
description: Защитные ограничения — это механизмы безопасности, размещаемые между пользователями, инструментами и моделями генеративного ИИ для оценки промптов, извлеченного контекста, выходных данных модели и действий агента до...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - AI Model Engineering
    - AI Model Evaluation
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-07-31"
source_name: Generative AI Guardrails
technique_count: 23
title: Защитные ограничения (Guardrails) для генеративного ИИ
url: /mitigations/AML.M0020/
---

Защитные ограничения — это механизмы безопасности, размещаемые между пользователями, инструментами и моделями генеративного ИИ для оценки промптов, извлеченного контекста, выходных данных модели и действий агента до того, как они будут приняты, выполнены или показаны пользователю. Они помогают блокировать, изменять или перенаправлять нежелательное содержимое, такое как вредоносный код, вредоносные инструкции, чувствительные данные, неподтвержденные утверждения, ответы, нарушающие политики, или небезопасные запросы к инструментам.

Защитные ограничения могут быть реализованы с помощью механизмов на основе правил, таких как фильтры, списки разрешений, списки блокировок, регулярные выражения, валидация схем, правила политик и проверки разрешений, либо с помощью техник на основе ИИ, таких как классификаторы, LLM-рецензенты или LLM-судьи, распознавание именованных сущностей (NER), проверки опоры на источники и проверки соответствия задаче. Они могут применяться на нескольких этапах рабочего процесса генеративного ИИ, включая обработку входных данных, формирование промпта, извлечение данных, выполнение инструментов, проверку выходных данных модели и мониторинг после развертывания.

Примеры конкретных реализаций защитных ограничений включают:[[owasp-llm-top10]]  [[datadog-llm-guardrails]] [[azure-ai-content-safety]] [[nvidia-nemo-guardrails]]
- Модерация входных данных: проверка пользовательских промптов на вредоносное содержимое, попытки prompt injection, попытки jailbreak, чувствительные данные, запросы не по теме или входные данные, превышающие ожидаемую длину либо не соответствующие ожидаемому формату.
- Модерация выходных данных: сканирование ответов модели перед отправкой пользователям на предмет вредоносного содержимого, PII, секретов, нарушений политик, неподтвержденных утверждений или небезопасного кода с помощью классических сканеров, классификаторов или выделенной модели-рецензента.
- Обеспечение соблюдения системного промпта и политик: обеспечение выполнения системных инструкций, ролей пользователей, границ предметной области, форматов ответов и политик отказа до того, как модель сформирует ответ (см. [Инструкции для генеративного ИИ](/mitigations/AML.M0021)).
- Защитные ограничения для инструментов и действий: проверка вызовов инструментов, аргументов инструментов, разрешений и выходных данных инструментов перед выполнением либо до возврата результатов модели. Требуйте подтверждения человеком для высокозначимых, необратимых, привилегированных или внешне видимых действий (см. [Участие человека в действиях ИИ-агента](/mitigations/AML.M0029), [Валидация входных и выходных данных компонентов ИИ-агента](/mitigations/AML.M0033)).
- Защитные ограничения при извлечении данных: фильтрация и проверка извлеченных документов перед добавлением в контекст модели, включая проверки на недоверенные источники, вредоносные инструкции, нерелевантный контекст или чувствительные данные.
- Проверки опоры на источники и фактической корректности: сравнение ответов модели с доверенными исходными материалами или утвержденными базами знаний для выявления неподтвержденных или галлюцинированных утверждений.
- Защита чувствительных данных и секретов: обнаружение, редактирование или блокирование персональной информации, учетных данных, токенов, проприетарных данных, системных промптов и другой конфиденциальной информации в промптах, извлеченном контексте, выходных данных инструментов и ответах модели.
- Валидация структурированных выходных данных: обеспечение соблюдения схем, проверок типов, разрешенных значений и безопасных форматов до того, как выходные данные модели будут использованы нижестоящими системами.

Защитные ограничения следует непрерывно оценивать, проверять с помощью red teaming и обновлять по мере развития состязательных техник. Решения защитных ограничений должны журналироваться (см. [Логирование телеметрии ИИ](/mitigations/AML.M0024)), а выявленные сбои следует систематически учитывать при обновлении политик, наборов данных для оценки, логики обнаружения, промптов и [Выравнивания модели генеративного ИИ](/mitigations/AML.M0022).


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong><p>Защитные ограничения могут обнаруживать вредоносный код в выходных данных модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><p>Apply guardrails to block malicious use and jailbreak attempts against generative AI models or services.</p></a>
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong><p>Защитные ограничения могут предотвращать вредоносные входные данные, способные привести к промпт-инъекции.</p></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong><p>Защитные ограничения могут предотвращать вредоносные входные данные, способные привести к компрометации плагина, и обнаруживать персональные данные в выходных данных модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0054/"><span class="relation-id">AML.T0054</span><strong>Джейлбрейк LLM</strong><p>Защитные ограничения могут предотвращать вредоносные входные данные, способные привести к джейлбрейку.</p></a>
<a class="relation-item" href="/techniques/AML.T0056/"><span class="relation-id">AML.T0056</span><strong>Извлечение системного промпта LLM</strong><p>Защитные ограничения могут предотвращать вредоносные входные данные, способные привести к извлечению метапромпта.</p></a>
<a class="relation-item" href="/techniques/AML.T0057/"><span class="relation-id">AML.T0057</span><strong>Утечка данных из LLM</strong><p>Защитные ограничения могут обнаруживать чувствительные данные и персональные данные в выходных данных модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0061/"><span class="relation-id">AML.T0061</span><strong>Саморепликация промпта LLM</strong><p>Защитные ограничения могут помогать предотвращать атаки репликации во входных и выходных данных модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0062/"><span class="relation-id">AML.T0062</span><strong>Выявление галлюцинированных сущностей LLM</strong><p>Защитные ограничения могут помогать блокировать галлюцинированный контент в выходных данных модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0066/"><span class="relation-id">AML.T0066</span><strong>Подготовка содержимого для извлечения</strong><p>Apply retrieval guardrails to reject untrusted, malicious, irrelevant, or unsupported RAG content.</p></a>
<a class="relation-item" href="/techniques/AML.T0068/"><span class="relation-id">AML.T0068</span><strong>Обфускация промпта LLM</strong><p>Apply input guardrails that decode, normalize, inspect, and block concealed malicious instructions.</p></a>
<a class="relation-item" href="/techniques/AML.T0069/"><span class="relation-id">AML.T0069</span><strong>Выявление системной информации LLM</strong><p>Block output of system prompts, hidden instructions, secrets, and internal configuration details.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><p>Filter responses that expose hidden prompt formatting, delimiters, or internal instructions.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.001/"><span class="relation-id">AML.T0069.001</span><strong>Ключевые слова системных инструкций</strong><p>Filter responses that reveal system keywords, tool names, function definitions, or hidden instructions.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><p>Filter responses that reveal system prompts and hidden instructions.</p></a>
<a class="relation-item" href="/techniques/AML.T0070/"><span class="relation-id">AML.T0070</span><strong>Отравление RAG</strong><p>Apply retrieval guardrails to reject untrusted, malicious, irrelevant, or unsupported RAG content.</p></a>
<a class="relation-item" href="/techniques/AML.T0071/"><span class="relation-id">AML.T0071</span><strong>Внедрение ложной записи RAG</strong><p>Validate retrieved documents and block nested false entries, malicious instructions, and unsupported content.</p></a>
<a class="relation-item" href="/techniques/AML.T0078/"><span class="relation-id">AML.T0078</span><strong>Компрометация при посещении сайта</strong><p>Apply retrieval, input, and tool-action guardrails to detect and block malicious instructions embedded in websites before execution.</p></a>
<a class="relation-item" href="/techniques/AML.T0099/"><span class="relation-id">AML.T0099</span><strong>Отравление данных инструмента ИИ-агента</strong><p>Apply retrieval guardrails to reject untrusted, malicious, irrelevant, or unsupported tool-retrieved content.</p></a>
<a class="relation-item" href="/techniques/AML.T0100/"><span class="relation-id">AML.T0100</span><strong>Кликбейт для ИИ-агента</strong><p>Apply guardrails to untrusted web content and proposed computer-use actions.</p></a>
<a class="relation-item" href="/techniques/AML.T0102/"><span class="relation-id">AML.T0102</span><strong>Генерация вредоносных команд</strong><p>Block prompts and outputs that request or contain malicious commands.</p></a>
<a class="relation-item" href="/techniques/AML.T0108/"><span class="relation-id">AML.T0108</span><strong>ИИ-агент</strong><p>Block malicious agent instructions, command outputs, and unsafe tool requests with guardrails.</p></a>
<a class="relation-item" href="/techniques/AML.T0112.000/"><span class="relation-id">AML.T0112.000</span><strong>Локальный ИИ-агент</strong><p>Block malicious prompts, unsafe outputs, and prohibited agent actions with guardrails.</p></a>
</div>
