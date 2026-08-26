---
atlas_id: AML.M0036
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2026-07-31"
description: Ограничивайте объём ресурсов, который может потреблять запрос к ИИ, задание инференса или рабочий процесс ИИ-агента. Устанавливайте ограничения на размер входных данных, размер батча, время выполнения, объём памяти,...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-07-31"
source_name: Limit AI Workload Resource Consumption
technique_count: 4
title: Ограничение потребления ресурсов рабочими нагрузками ИИ
url: /mitigations/AML.M0036/
---

Ограничивайте объём ресурсов, который может потреблять запрос к ИИ, задание инференса или рабочий процесс ИИ-агента. Устанавливайте ограничения на размер входных данных, размер батча, время выполнения, объём памяти, вычислительные ресурсы и размер выходных данных. Сервисы генеративного ИИ также должны ограничивать число токенов контекста и выходных токенов. В агентных ИИ-системах следует ограничивать число итераций, повторных попыток, вызовов инструментов и параллельных задач, глубину делегирования и расходы, возникающие при обращении к нижестоящим сервисам.

Применяйте ограничения ресурсов на всём протяжении рабочего процесса, поскольку один запрос может инициировать несколько вызовов модели или внешних действий. Используйте таймауты, предельные уровни затрат, автоматические выключатели (circuit breakers) и условия безопасного завершения, чтобы отдельные рабочие нагрузки не могли исчерпать общие ресурсы. Эти меры дополняют ограничения частоты запросов, которые защищают от большого числа запросов, по отдельности не требующих значительных затрат.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0029/"><span class="relation-id">AML.T0029</span><strong>Отказ в обслуживании ИИ-сервиса</strong><p>Limit the resources consumed by individual requests to reduce denial of service from computationally expensive inputs.</p></a>
<a class="relation-item" href="/techniques/AML.T0034/"><span class="relation-id">AML.T0034</span><strong>Искусственное увеличение затрат</strong><p>Apply resource budgets to AI requests and workflows to limit costs from resource-intensive queries, excessive output, and uncontrolled agentic activity.</p></a>
<a class="relation-item" href="/techniques/AML.T0034.001/"><span class="relation-id">AML.T0034.001</span><strong>Ресурсоёмкие запросы</strong><p>Bound input size, output size, execution time, memory, and compute consumed by resource-intensive queries.</p></a>
<a class="relation-item" href="/techniques/AML.T0034.002/"><span class="relation-id">AML.T0034.002</span><strong>Потребление ресурсов агентом</strong><p>Limit agent iterations, tool calls, fan-out, runtime, and downstream spending to constrain agentic resource consumption.</p></a>
</div>
