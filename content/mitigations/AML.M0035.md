---
atlas_id: AML.M0035
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Policy
    - Technical - AI
    - Technical - Cyber
created_date: "2026-07-21"
description: Создайте красную команду по ИИ, ответственную за проведение регулярных, санкционированных учений красной команды, основанных на данных об угрозах, для выявления и устранения уязвимостей в системах с поддержкой ИИ до...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
    - Data Preparation
    - AI Model Engineering
    - AI Model Evaluation
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-07-21"
source_name: AI Red Team
technique_count: 33
title: Красная команда по ИИ
url: /mitigations/AML.M0035/
---

Создайте красную команду по ИИ, ответственную за проведение регулярных, санкционированных учений красной команды, основанных на данных об угрозах, для выявления и устранения уязвимостей в системах с поддержкой ИИ до развёртывания и на всём протяжении эксплуатации. Тестирование ИИ силами красной команды имитирует реалистичное поведение злоумышленников, чтобы оценить, как атаки могут повлиять на конфиденциальность, целостность, доступность, безопасность, приватность и эффективность выполнения целевых задач системой с поддержкой ИИ.

Учения красной команды должны охватывать всю систему с поддержкой ИИ, включая модели и данные, агентов (в том числе их память и инструменты), потоки данных, процессы принятия решений, логику приложений, системы извлечения данных, идентичности и разрешения, зависимости ПО, компоненты системы, не использующие ИИ, инфраструктуру, пользовательские интерфейсы и рабочие процессы с участием людей.

Учение красной команды по ИИ можно организовать в три этапа: планирование и определение области учения, проведение выбранных испытаний и оценка результатов, на основе которой формируются отчётность и меры по устранению недостатков.

1. **Планирование и определение области**
    - Документируйте предполагаемое использование системы, среду развёртывания, пользователей, чувствительные данные, подключённые ресурсы и возможные последствия отказа или неправомерного использования. Составьте схему компонентов системы, границ доверия, потоков данных, внешних сервисов, точек принятия решений человеком, а также точек доступа на этапах обучения и инференса.
    - Установите правила проведения учений, охватывающие разрешённые системы, учётные записи, данные и техники, временные окна тестирования, ограничения ресурсов, процедуры эскалации, порядок работы с доказательствами и условия остановки. Планируйте разрушительные, нарушающие приватность или дорогостоящие тесты для изолированных сред с надлежащими мерами защиты.
    - Разработайте модель угроз на основе среды эксплуатации системы и релевантного поведения злоумышленников. Определите цели злоумышленника, уровень его доступа, знания, возможности, ресурсы и ограничения; учтите цифровые и физические пути атак, а также роль контроля со стороны человека.
    - Используйте тактики, техники и процедуры ATLAS для выявления релевантных вариантов поведения злоумышленников и построения векторов угроз. Назначьте им приоритеты с учётом вероятности и серьёзности воздействия на систему.
    - Определите критерии успеха и условия остановки. Определите метрики уровня задач для оценки воздействия на возможности ИИ и операционные метрики для измерения воздействия на систему в целом.

2. **Проведение**
    - Проводите выбранные испытания ручными и автоматизированными методами по мере необходимости. Автоматизация позволяет создавать варианты входных данных, повторно воспроизводить последовательности атак и оценивать ответы в большом масштабе. Специалисты, проводящие тестирование, могут разрабатывать атаки с учётом особенностей системы, адаптироваться к наблюдаемым средствам защиты и исследовать неожиданное поведение.
    - Соблюдайте правила проведения учений и фиксируйте действия в рамках атаки, ответы системы, поведение средств контроля, отклонения от плана тестирования и доказательства, необходимые для оценки результатов.
    - Остановите тестирование или запустите процедуру эскалации при достижении заранее определённых условий. После тестирования удалите тестовые учётные записи, изменённые данные, установленное ПО, инструкции, сохраняющиеся в системе, и другие артефакты учений.

3. **Оценка, отчётность и совершенствование**
    - Оцените результаты по заданным метрикам уровня задач и операционным метрикам. Документируйте успешные и неуспешные атаки, их последствия, наблюдаемое поведение средств контроля, отклонения от плана тестирования и пробелы в модели угроз.
    - Сообщите о выявленных проблемах соответствующим разработчикам, специалистам по защите, командам эксплуатации, владельцам рисков и другим заинтересованным сторонам.
    - Назначьте ответственных за выявленные проблемы, отслеживайте их устранение и повторно тестируйте исправленные системы.
    - Используйте продемонстрированные атаки для совершенствования превентивных средств контроля, обнаружения, реагирования на инциденты и восстановления. Там, где это уместно, преобразуйте подтверждённые сбои в регрессионные тесты, наборы данных для оценки, логику обнаружения, требования к мониторингу или критерии развёртывания.

Тестирование силами красной команды — непрерывный процесс; его следует повторять по мере развития ландшафта угроз, а также при внесении изменений в систему, её компоненты, предполагаемое использование или среду развёртывания. Новые данные об угрозах, уязвимости и результаты тестирования должны служить основой для определения области и приоритетов будущих учений.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong><p>Отработайте в контролируемых условиях внедрение недоверенных компонентов — ПО, данных, моделей и инструментов ИИ-агента — через репрезентативные пути их получения и развёртывания. Устраните недостатки, связанные со сведениями о происхождении, валидацией, процедурами одобрения, изоляцией и откатом.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><p>Внедряйте недоверенные пакеты ПО для ИИ, библиотеки, плагины или программные компоненты только в контролируемую среду. Проверьте средства контроля зависимостей, сканирование, подписание, процедуры одобрения, изоляцию и безопасность установки.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><p>Внедряйте недоверенные наборы данных только в контролируемую среду через репрезентативные пути их получения и приёма. Проверьте механизмы отслеживания происхождения, обеспечения целостности, санитизации, проверки и отбраковки данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><p>Внедряйте недоверенную или модифицированную модель только в контролируемую среду через репрезентативные пути её получения и развёртывания. Проверьте сведения о происхождении, сканирование, подписание, процедуры одобрения, изоляцию и возможность возврата к предыдущей версии модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><p>Внедряйте недоверенный инструмент ИИ-агента или недоверенное определение инструмента только в контролируемую среду. Убедитесь, что источник входит в число разрешённых; проверьте целостность, факт проведения проверки, границы разрешений и безопасность активации.</p></a>
<a class="relation-item" href="/techniques/AML.T0015/"><span class="relation-id">AML.T0015</span><strong>Обход ИИ-модели</strong><p>Проведите репрезентативные цифровые и мультимодальные атаки обхода, а также атаки обхода в физическом домене. Используйте успешные атаки, чтобы повысить устойчивость модели и улучшить предобработку, обнаружение состязательных входных данных, контроль со стороны человека и мониторинг.</p></a>
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong><p>В контролируемых условиях попытайтесь изменить или подменить модели, веса, адаптеры и связанную с ними конфигурацию. Устраните недостатки механизмов авторизации, обеспечения целостности артефактов, процедур одобрения развёртывания, мониторинга и восстановления.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><p>Проверьте, могут ли контролируемые изменения весов модели, контролируемое дообучение или контролируемые изменения связанных с моделью артефактов привести к появлению целевого или устойчивого поведения. Усовершенствуйте отслеживание происхождения модели, её валидацию и мониторинг целостности, а также механизм возврата к предыдущей версии.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><p>В контролируемых условиях попытайтесь внести неавторизованные изменения в архитектуру модели или в её исполняемые компоненты. Проверьте процессы рассмотрения изменений, контроля целостности, подписания и одобрения развёртывания, а также механизмы восстановления.</p></a>
<a class="relation-item" href="/techniques/AML.T0020/"><span class="relation-id">AML.T0020</span><strong>Отравление обучающих данных</strong><p>Introduce controlled poisoned records or triggers into representative data pipelines. Verify and improve provenance, sanitization, review, drift detection, model validation, and rollback controls.</p></a>
<a class="relation-item" href="/techniques/AML.T0024/"><span class="relation-id">AML.T0024</span><strong>Эксфильтрация через API инференса ИИ</strong><p>Exercise inference interfaces for membership inference, model inversion, and functional extraction. Use findings to improve privacy controls, authentication, output restriction, rate limits, and monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.000/"><span class="relation-id">AML.T0024.000</span><strong>Определение принадлежности к обучающей выборке</strong><p>Test whether model outputs reveal the membership of known training samples. Improve privacy-preserving training, output restriction, access controls, and query monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><p>Attempt to reconstruct sensitive records, attributes, or representative training information from model outputs. Remediate leakage through model changes, output minimization, privacy controls, and restricted inference access.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><p>Simulate functional model extraction through inference queries. Establish appropriate authentication, rate limits, output restrictions, anomaly detection, and extraction monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0029/"><span class="relation-id">AML.T0029</span><strong>Отказ в обслуживании ИИ-сервиса</strong><p>Submit adversarial workloads and exercise dependency failures that could exhaust inference or supporting services. Apply quotas, concurrency limits, timeouts, resource isolation, and graceful degradation based on findings.</p></a>
<a class="relation-item" href="/techniques/AML.T0034/"><span class="relation-id">AML.T0034</span><strong>Искусственное увеличение затрат</strong><p>Exercise requests and workflows designed to amplify inference, infrastructure, or external-service costs. Verify budgets, quotas, rate limits, workload controls, alerting, and termination mechanisms.</p></a>
<a class="relation-item" href="/techniques/AML.T0034.000/"><span class="relation-id">AML.T0034.000</span><strong>Чрезмерные запросы</strong><p>Generate controlled high-volume query activity. Verify authentication, user and tenant quotas, rate limits, anomaly detection, cost alerts, and service protection.</p></a>
<a class="relation-item" href="/techniques/AML.T0034.001/"><span class="relation-id">AML.T0034.001</span><strong>Ресурсоёмкие запросы</strong><p>Submit controlled requests designed to consume disproportionate inference resources. Verify input constraints, timeouts, workload limits, resource isolation, and cost monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0034.002/"><span class="relation-id">AML.T0034.002</span><strong>Потребление ресурсов агентом</strong><p>Test recursive behavior, repeated tool calls, costly API use, and attacker-controlled task expansion. Verify budgets, iteration limits, timeouts, approval thresholds, and termination controls.</p></a>
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong><p>Test direct, indirect, and triggered instructions through user input, retrieved data, documents, messages, websites, images, metadata, and tool output. Remediate trust-boundary, instruction-handling, permission, and monitoring failures.</p></a>
<a class="relation-item" href="/techniques/AML.T0051.000/"><span class="relation-id">AML.T0051.000</span><strong>Прямая промпт-инъекция</strong><p>Submit controlled malicious instructions directly through user-facing model and application interfaces. Improve instruction enforcement, input controls, authorization, and monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0051.001/"><span class="relation-id">AML.T0051.001</span><strong>Косвенная промпт-инъекция</strong><p>Place controlled malicious instructions in external or retrieved content processed by the system. Improve content trust boundaries, retrieval controls, instruction isolation, and restrictions on resulting actions.</p></a>
<a class="relation-item" href="/techniques/AML.T0051.002/"><span class="relation-id">AML.T0051.002</span><strong>Триггерная промпт-инъекция</strong><p>Place controlled instructions in content or workflows where later user actions or system events activate them. Verify trigger authorization, context handling, action restrictions, and monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong><p>Attempt to select unauthorized tools, supply unsafe arguments, exceed user privileges, or chain tools into harmful actions. Correct permissions, argument validation, sandboxing, action controls, and approval requirements.</p></a>
<a class="relation-item" href="/techniques/AML.T0054/"><span class="relation-id">AML.T0054</span><strong>Джейлбрейк LLM</strong><p>Exercise manual and automated multi-turn, multilingual, encoded, transformed, and multimodal jailbreaks. Incorporate successful cases into guardrails, guidelines, alignment, monitoring, and regression evaluations.</p></a>
<a class="relation-item" href="/techniques/AML.T0056/"><span class="relation-id">AML.T0056</span><strong>Извлечение системного промпта LLM</strong><p>Probe model and application interfaces for disclosure of system instructions, policies, tool definitions, or hidden context. Remove embedded secrets and improve configuration isolation and output controls.</p></a>
<a class="relation-item" href="/techniques/AML.T0057/"><span class="relation-id">AML.T0057</span><strong>Утечка данных из LLM</strong><p>Place synthetic secrets or canary records in representative data sources and attempt extraction through prompts, retrieval, tools, and rendered output. Improve authorization boundaries, filtering, tenant isolation, and exfiltration detection.</p></a>
<a class="relation-item" href="/techniques/AML.T0068/"><span class="relation-id">AML.T0068</span><strong>Обфускация промпта LLM</strong><p>Test encoded, transformed, visually hidden, multilingual, and multimodal instructions. Use successful bypasses to improve normalization, decoding, content inspection, and detection controls.</p></a>
<a class="relation-item" href="/techniques/AML.T0070/"><span class="relation-id">AML.T0070</span><strong>Отравление RAG</strong><p>Seed controlled malicious or misleading content into representative ingestion sources. Improve source authorization, provenance, content validation, indexing controls, and retrieval-time filtering.</p></a>
<a class="relation-item" href="/techniques/AML.T0080/"><span class="relation-id">AML.T0080</span><strong>Отравление контекста ИИ-агента</strong><p>Attempt to persist malicious instructions in agent memory and long-lived threads. Verify authorization for context changes, integrity checks, trust labeling, expiration, user visibility, and remediation.</p></a>
<a class="relation-item" href="/techniques/AML.T0080.000/"><span class="relation-id">AML.T0080.000</span><strong>Память</strong><p>Attempt to store controlled malicious instructions or preferences in persistent agent memory. Verify authorization, user visibility, integrity validation, expiration, and memory-remediation controls.</p></a>
<a class="relation-item" href="/techniques/AML.T0080.001/"><span class="relation-id">AML.T0080.001</span><strong>Цепочка сообщений</strong><p>Introduce controlled malicious instructions into long-lived or shared conversation threads. Verify context isolation, trust handling, thread reset, expiration, and monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0081/"><span class="relation-id">AML.T0081</span><strong>Изменение конфигурации ИИ-агента</strong><p>Exercise unauthorized changes to system prompts, tools, knowledge sources, security settings, and approval requirements. Improve access controls, change approval, integrity monitoring, and restoration from trusted configurations.</p></a>
</div>
