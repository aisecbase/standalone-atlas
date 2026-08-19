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
description: Establish an AI red team responsible for conducting recurring, authorized, and threat-informed red-teaming exercises to identify and remediate vulnerabilities in AI-enabled systems before deployment and throughout...
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
title: AI Red Team
url: /mitigations/AML.M0035/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Establish an AI red team responsible for conducting recurring, authorized, and threat-informed red-teaming exercises to identify and remediate vulnerabilities in AI-enabled systems before deployment and throughout operation. AI red-teaming simulates realistic adversary behavior to evaluate how attacks could affect the confidentiality, integrity, availability, safety, privacy, and mission performance of an AI-enabled system.

Red-teaming exercises should consider the complete AI-enabled system, including models and data, agents (including memory and tools), data flows, decision processes, application logic, retrieval systems, identities and permissions, software dependencies, non-AI system components, infrastructure, user interfaces, and human workflows.

An AI red team exercise can be organized into three phases: planning and scoping the exercise, executing the selected exercises, and assessing the results to guide reporting and remediation.

1. **Plan and Scope**
    - Document the system's intended use, deployment environment, users, sensitive data, connected resources, and potential consequences of failure or misuse. Diagram the system's components, trust boundaries, data flows, external services, human decision points, and training- and inference-time access points.
    - Establish rules of engagement covering authorized systems, accounts, data, techniques, test windows, resource limits, escalation procedures, evidence handling, and stop conditions. Plan destructive, privacy-invasive, or high-cost tests for isolated environments with appropriate safeguards.
    - Develop a threat model based on the system's operating environment and relevant adversary behavior. Define the adversary's objectives, access, knowledge, capabilities, resources, and constraints, and account for digital and physical attack paths and the role of human oversight.
    - Use ATLAS tactics, techniques, and procedures to identify relevant adversary behaviors and construct threat vectors. Prioritize them according to likelihood and severity of impact to the system.
    - Define success criteria and stopping conditions. Identify task-level metrics for effects on the AI capability and operational metrics for measuring impact to the overall system.

2. **Execute**
    - Conduct the selected exercises using manual and automated methods as appropriate. Automation can generate input variations, replay attack sequences, and evaluate responses at scale. Human testers can develop system-specific attacks, adapt to observed defenses, and investigate unexpected behavior.
    - Follow the rules of engagement and record attack activity, system responses, control behavior, deviations from the test plan, and evidence needed to evaluate the results.
    - Stop or escalate testing when predefined conditions are reached. After testing, remove test accounts, modified data, installed software, persistent instructions, and other exercise artifacts.

3. **Assess, Report, and Improve**
    - Evaluate the results against the defined task and operational metrics. Document successful and unsuccessful attacks, their consequences, observed control behavior, deviations from the test plan, and gaps in the threat model.
    - Report findings to the appropriate developers, defenders, operational teams, risk owners, and other stakeholders.
    - Assign findings to responsible owners, track remediation, and retest corrected systems.
    - Use demonstrated attacks to improve preventive controls, detection, incident response, and recovery. Where appropriate, convert confirmed failures into regression tests, evaluation datasets, detection logic, monitoring requirements, or deployment criteria.

Red-teaming is a continuous process and should be repeated as the threat landscape evolves and when changes are made to the system, its components, intended use, or deployment environment. New threat intelligence, vulnerabilities, and test results should inform the scope and priorities of future exercises.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong><p>Exercise the introduction of controlled untrusted software, data, models, and agent tools through representative acquisition and deployment paths. Remediate weaknesses in provenance, validation, approvals, isolation, and rollback.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><p>Introduce controlled untrusted AI packages, libraries, plugins, or software components. Verify dependency controls, scanning, signing, approval, isolation, and safe installation.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><p>Introduce controlled untrusted datasets through representative acquisition and ingestion paths. Verify provenance, integrity, sanitization, review, and rejection controls.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><p>Introduce a controlled untrusted or modified model through representative acquisition and deployment paths. Verify provenance, scanning, signing, approval, isolation, and rollback.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><p>Introduce a controlled untrusted agent tool or tool definition. Verify source authorization, integrity, review, permission boundaries, and safe activation.</p></a>
<a class="relation-item" href="/techniques/AML.T0015/"><span class="relation-id">AML.T0015</span><strong>Обход ИИ-модели</strong><p>Execute representative digital, multimodal, and physical-domain evasion attacks. Use successful attacks to improve model robustness, preprocessing, adversarial-input detection, human oversight, and monitoring.</p></a>
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong><p>Attempt controlled modification or substitution of models, weights, adapters, and related configuration. Remediate weaknesses in authorization, artifact integrity, deployment approval, monitoring, and recovery.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><p>Test whether controlled changes to model weights, fine-tuning, or associated artifacts can introduce targeted or persistent behavior. Improve model provenance, validation, integrity monitoring, and rollback.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><p>Attempt controlled unauthorized changes to model architecture or executable model components. Verify review, integrity checking, signing, deployment approval, and restoration controls.</p></a>
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
