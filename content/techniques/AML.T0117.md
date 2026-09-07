---
atlas_id: AML.T0117
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Злоумышленники могут использовать ИИ-агента для автономного построения и многократного пересмотра пути атаки, ведущего к заданной злоумышленником цели. Получив цель высокого уровня, система может формировать...
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
procedure_count: 4
source_name: Autonomous Attack-Path Adaptation
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Автономная адаптация пути атаки
url: /techniques/AML.T0117/
---

Злоумышленники могут использовать ИИ-агента для автономного построения и многократного пересмотра пути атаки, ведущего к заданной злоумышленником цели. Получив цель высокого уровня, система может формировать промежуточные цели, определять необходимые условия, сравнивать возможные пути и учитывать наблюдения, полученные в результате предыдущих действий, чтобы адаптивно выстраивать последовательность техник без руководства человеком на каждом шаге.

Автономные ИИ-агенты могут проявлять такое поведение и при выполнении задачи, поставленной в законных, безвредных или разрешённых целях, если выбранные системой промежуточные цели или методы пересекают границы полномочий, доверия, контроля или безопасности и приводят к попытке осуществить вредоносную киберактивность либо к её фактическому осуществлению.

В ходе повторяющихся циклов «наблюдение — решение — действие» система может интерпретировать выходные данные команд, ошибки, реакцию защитных средств, изменения в доступе и вновь обнаруженные сведения. На основе этих наблюдений она может менять приоритет действий, заменять промежуточную цель, отказываться от непродуктивной ветви, отбрасывать выводы, опровергнутые дополнительными доказательствами, или следовать альтернативным путём.

Автономная ИИ-система может формировать вспомогательные цели, основное назначение которых — расширить её будущие операционные возможности, а не непосредственно выполнить поставленную задачу. Такие цели могут включать получение новых эксплойтов (см. [Автономная разработка эксплойтов](/techniques/AML.T0017.001)), дополнительных полномочий, идентичностей, сред выполнения, каналов связи, инструментов или отношений доверия, расширяющих набор действий, доступных в последующих циклах планирования. Вновь полученные возможности сами могут стать необходимыми условиями для новых вспомогательных целей, что приводит к постепенному расширению области действий агента в ходе операции.

Перепланирование пути атаки может происходить в рамках одного запуска агента или возникать при работе нескольких независимых агентов. Агенты могут обмениваться информацией посредством сохраняемых общих артефактов (см. [Коммуникация автономных ИИ-агентов: коммуникация через общие артефакты](/techniques/AML.T0118.000)), благодаря чему обнаруженные сведения, запросы, возможности, ограничения, состояние задачи и результаты одного агента могут влиять на последующий путь, выбранный другим. Участвующие агенты могут принимать к выполнению запросы друг друга, добровольно распределять работу, повторно использовать успешные методы, продолжать незавершённые действия или перенаправлять собственные локальные пути без централизованного планировщика, общего окна контекста или полного представления об операции в целом.

Участие человека не исключает автономного перепланирования пути атаки. Оператор, пользователь, специалист по оценке или рабочий процесс могут выбирать цель атаки, определять задачу, устанавливать ограничения, предоставлять возможности или одобрять переходы, влекущие существенные последствия.

[Автономная адаптация пути атаки](/techniques/AML.T0117) и [Автономная оркестрация атаки](/techniques/AML.T0124) могут применяться совместно, однако описывают разные функции управления. Адаптация пути атаки отражает то, как полученные свидетельства изменяют выбранный путь. Оркестрация атаки отражает то, как работа распределяется между агентами, координируется, проверяется и перенаправляется.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0037/"><span class="relation-id">AML.M0037</span><strong>Контроль расширения полномочий ИИ-агента</strong><p>Если организация обладает достаточным уровнем административного контроля над ИИ-системой, чтобы обеспечивать соблюдение установленных для неё границ полномочий, контроль расширения полномочий ИИ-агента может непосредственно ограничивать достижение вспомогательных целей, направленных на получение новых полномочий, идентичностей, сред выполнения, инструментов, каналов связи или отношений доверия. Эти средства контроля не ограничивают подконтрольные злоумышленнику ИИ-системы, над которыми у организации нет административного контроля.</p></a>
<a class="relation-item" href="/mitigations/AML.M0038/"><span class="relation-id">AML.M0038</span><strong>Выявление отклонений ИИ-агента от установленной области действий</strong><p>Если организация обладает достаточным уровнем административного контроля над ИИ-системой, чтобы отслеживать её планы и действия, выявление отклонений ИИ-агента от установленной области действий позволяет проверить, остаются ли динамически формируемые промежуточные действия в рамках поставленной агенту задачи. Эта мера неприменима к подконтрольным злоумышленнику ИИ-системам, над которыми у организации нет административного контроля.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Перед агентами стояла цель выполнить задания ExploitGym. Исчерпав предусмотренные способы решения, они начали исследовать доступное им окружение и разрабатывать альтернативные способы выполнения своих задач. Агенты формировали промежуточные цели и неоднократно перестраивали путь атаки, включавший обход изоляции, использование внешней инфраструктуры, получение материалов, связанных с заданием, эксплуатацию уязвимостей Hugging Face, получение доступа к учётным данным и сбор данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>GTG-1002 assigned their Claude agent target-scoped objectives against a human-selected organization under false defensive-testing context. Between operator-controlled stage gates, the agent derived and revised intermediate actions for reconnaissance, vulnerability exploitation, credential access, internal navigation, collection, and exfiltration, selecting and invoking available tools based on operational results.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>After receiving an initial task, DeepSeek sequenced reconnaissance and exploitation actions, evaluated failed prerequisites, abandoned Langflow, compared alternative products and vulnerabilities, and selected n8n. Unit 42 recovered no additional operator input during the session.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The framework constructed numerous candidate multi-step attack paths using confirmed prerequisites, observed blockers, and estimated success probabilities. It promoted paths supported by validated evidence, queued paths requiring additional investigation, discarded false positives and blocked paths, and initiated target-specific learning cycles when existing methods failed.</p></a>
</div>
