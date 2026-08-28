---
atlas_id: AML.T0051
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленник может сформировать вредоносные промпты как входные данные для LLM, чтобы заставить модель действовать непреднамеренным образом. Такие «промпт-инъекции» часто предназначены для того, чтобы модель...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 7
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: LLM Prompt Injection
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0005
title: Промпт-инъекция в LLM
url: /techniques/AML.T0051/
---

Злоумышленник может сформировать вредоносные промпты как входные данные для LLM, чтобы заставить модель действовать непреднамеренным образом.

Такие «промпт-инъекции» часто предназначены для того, чтобы модель игнорировала части своих исходных инструкций и вместо этого следовала инструкциям злоумышленника.

Промпт-инъекция может быть вектором первичного доступа к LLM и дать злоумышленнику точку закрепления для дальнейших этапов операции. Она может быть рассчитана на обход защитных механизмов LLM или на возможность отдавать привилегированные команды. Эффект промпт-инъекции может сохраняться на протяжении интерактивной сессии с LLM.

Вредоносные промпты могут внедряться злоумышленником напрямую ([прямой вариант](/techniques/AML.T0051.000)) для генерации вредоносного содержимого через LLM или для закрепления в системе и последующих воздействий. Промпты также могут внедряться косвенно, когда LLM в рамках своей обычной работы получает вредоносный промпт из другого источника данных ([косвенный вариант](/techniques/AML.T0051.001)). Такой тип инъекции может использоваться злоумышленником для закрепления в системе или для атаки на пользователя LLM. Вредоносные промпты также могут срабатывать по [триггеру](/techniques/AML.T0051.002), например из-за действий пользователя или системных событий.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051.000/"><span class="relation-id">AML.T0051.000</span><strong>Прямая промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.001/"><span class="relation-id">AML.T0051.001</span><strong>Косвенная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.002/"><span class="relation-id">AML.T0051.002</span><strong>Триггерная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Используйте контроль доступа в продакшене, чтобы помешать злоумышленникам внедрять вредоносные промпты.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Гардрейлы могут предотвращать вредоносные входные данные, способные привести к промпт-инъекции.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Инструкции для модели могут предписывать ей отказываться отвечать на небезопасные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/mitigations/AML.M0033/"><span class="relation-id">AML.M0033</span><strong>Валидация входных и выходных данных компонентов ИИ-агента</strong><p>Валидация может помешать злоумышленникам выполнять промпт-инъекции, способные повлиять на агентные рабочие процессы.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Тестируйте прямые, косвенные и срабатывающие по триггеру инструкции, передаваемые через пользовательский ввод, извлечённые данные, документы, сообщения, веб-сайты, изображения, метаданные и выходные данные инструментов. Устраняйте недостатки в обеспечении границ доверия, обработке инструкций, управлении разрешениями и мониторинге.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0005 Выполнение</span><p>Исследователи изменяют исходный промпт, чтобы обнаружить другие источники знаний и инструменты, которые могут содержать интересующие их данные.</p></a>
</div>
