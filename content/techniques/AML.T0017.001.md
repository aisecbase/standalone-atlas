---
atlas_id: AML.T0017.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Автономный ИИ-агент может выявить уязвимость ПО и разработать или существенно адаптировать средство её эксплуатации, получая лишь ограниченные указания от человека. Агент может анализировать исходный код,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 6
source_name: Autonomous Exploit Development
subtechnique_count: 0
subtechnique_of: AML.T0017
tactics:
    - AML.TA0003
title: Автономная разработка эксплойтов
url: /techniques/AML.T0017.001/
---

Автономный ИИ-агент может выявить уязвимость ПО и разработать или существенно адаптировать средство её эксплуатации, получая лишь ограниченные указания от человека. Агент может анализировать исходный код, документацию, поведение сервиса и сообщения об ошибках, чтобы сделать вывод о наличии уязвимости и определить условия, необходимые для её эксплуатации.

Агент может формулировать и проверять гипотезы об уязвимостях, генерировать проверочные запросы или полезные нагрузки, интерпретировать результаты и корректировать свой подход в ходе повторяющихся циклов «действие — наблюдение». Он может объединить несколько слабых мест в цепочку эксплойтов либо подготовить созданное средство эксплуатации к последующему использованию. Проверка может подтвердить, что эксплойт обеспечивает требуемый доступ, выполнение кода или иной технический результат. Уязвимость может быть публично известной или ранее неизвестной.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017/"><span class="relation-id">AML.T0017</span><strong>Разработка средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017.000/"><span class="relation-id">AML.T0017.000</span><strong>Состязательные атаки на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017.002/"><span class="relation-id">AML.T0017.002</span><strong>AI Agent Tools</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>An agent tested Artifactory request handling, inspected responses, and iteratively developed and validated an SSRF exploit that caused the package cache to retrieve external content for the isolated evaluation environment. OpenAI and JFrog characterized the vulnerability as a previously unknown zero-day.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>An agent probed Artifactory and identified an unauthenticated WebDAV MKCOL directory-creation path in a remote cache. It confirmed that arbitrary directory names persisted and could be enumerated by other runs, producing a reusable method for reconstructing the shared message board.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The agents developed and validated two execution methods against the exposed harness: redefining sqlite3_initialize so a submitted library invoked shell commands, and injecting shell syntax into submission path metadata.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The agents developed and validated an HDF5 artifact configuration that caused a dataset-processing worker to treat local filesystem content as external dataset storage.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Using the disclosed source code and configuration-processing details, the agents developed and validated a ReferenceFileSystem and Jinja2 exploit that produced arbitrary Python execution in a dataset conversion worker.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0003 Подготовка ресурсов</span><p>GTG-1002&#39;s Claude agent researched exploitation techniques for the identified SSRF vulnerability, generated a tailored custom payload and full exploit chain, tested the approach, evaluated the results, and adapted it for the target.</p></a>
</div>
