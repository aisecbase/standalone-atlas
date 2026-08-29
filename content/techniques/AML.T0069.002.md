---
atlas_id: AML.T0069.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут выявлять системные инструкции большой языковой модели, заданные разработчиком ИИ-системы, чтобы узнать о возможностях системы и обойти её гардрейлы.
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: System Prompt
subtechnique_count: 0
subtechnique_of: AML.T0069
tactics:
    - AML.TA0008
title: Системный промпт
url: /techniques/AML.T0069.002/
---

Злоумышленники могут выявлять системные инструкции большой языковой модели, заданные разработчиком ИИ-системы, чтобы узнать о возможностях системы и обойти её гардрейлы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0069/"><span class="relation-id">AML.T0069</span><strong>Выявление системной информации LLM</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.001/"><span class="relation-id">AML.T0069.001</span><strong>Ключевые слова системных инструкций</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Ограничивайте доступ к хранящимся системным промптам и шаблонам промптов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Authenticate and monitor access to production models and prompt configuration.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Фильтруйте ответы, раскрывающие системные промпты и скрытые инструкции.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учётные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0008 Выявление</span><p>Исследователь попросил ClawdBot выполнить `cat SOUL.md`, где `SOUL.md` — файл с системным промптом ClawdBot; в ответ ClawdBot вернул содержимое файла.</p></a>
</div>
