---
atlas_id: AML.M0000
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Policy
created_date: "2023-04-12"
description: Ограничьте публичное раскрытие технической информации об ИИ-стеке, используемом в продуктах или сервисах организации. Технические сведения о том, как используется ИИ, могут быть использованы злоумышленниками для...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
modified_date: "2026-07-31"
source_name: Limit Public Release of Information
technique_count: 15
title: Ограничение публичного раскрытия информации
url: /mitigations/AML.M0000/
---

Ограничьте публичное раскрытие технической информации об ИИ-стеке, используемом в продуктах или сервисах организации.
Технические сведения о том, как используется ИИ, могут быть использованы злоумышленниками для выбора целей и адаптации атак к целевой системе.
Кроме того, рассмотрите возможность ограничить раскрытие организационной информации, включая физические адреса, имена исследователей и структуру подразделений, на основании которой можно вывести технические детали, такие как техники ИИ, архитектуры моделей или наборы данных.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000/"><span class="relation-id">AML.T0000</span><strong>Поиск в открытых технических базах данных</strong><p>Ограничьте связь между публично раскрытыми подходами и данными, моделями и алгоритмами, используемыми в продакшене.</p></a>
<a class="relation-item" href="/techniques/AML.T0001/"><span class="relation-id">AML.T0001</span><strong>Поиск открытых материалов по анализу уязвимостей ИИ</strong><p>Limit disclosure of the production AI stack and system-specific technical details that let adversaries connect public vulnerability research to the deployed target.</p></a>
<a class="relation-item" href="/techniques/AML.T0002/"><span class="relation-id">AML.T0002</span><strong>Получение публичных ИИ-артефактов</strong><p>Ограничьте публикацию чувствительной информации в метаданных развернутых систем и публично доступных приложений.</p></a>
<a class="relation-item" href="/techniques/AML.T0003/"><span class="relation-id">AML.T0003</span><strong>Поиск на сайтах организации-жертвы</strong><p>Ограничьте публикацию технической информации о продуктах с поддержкой ML и организационной информации о командах, сопровождающих такие продукты.</p></a>
<a class="relation-item" href="/techniques/AML.T0004/"><span class="relation-id">AML.T0004</span><strong>Поиск в репозиториях приложений</strong><p>Ограничьте публикацию чувствительной информации в метаданных развернутых систем и публично доступных приложений.</p></a>
<a class="relation-item" href="/techniques/AML.T0005/"><span class="relation-id">AML.T0005</span><strong>Создание прокси-модели ИИ</strong><p>Ограничение публикации технической информации о модели и обучающих данных может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0005.000/"><span class="relation-id">AML.T0005.000</span><strong>Обучение прокси-модели на собранных ИИ-артефактах</strong><p>Ограничение публикации технической информации о модели и обучающих данных может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0005.002/"><span class="relation-id">AML.T0005.002</span><strong>Использование предварительно обученной модели</strong><p>Ограничение публикации технической информации о модели и обучающих данных может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0064/"><span class="relation-id">AML.T0064</span><strong>Сбор целей, индексируемых RAG</strong><p>Withhold public documentation that identifies RAG data sources, indexes, and retrieval architecture.</p></a>
<a class="relation-item" href="/techniques/AML.T0069/"><span class="relation-id">AML.T0069</span><strong>Выявление системной информации LLM</strong><p>Withhold public prompt templates, model configuration, and architecture details that could aid LLM reconnaissance.</p></a>
<a class="relation-item" href="/techniques/AML.T0084/"><span class="relation-id">AML.T0084</span><strong>Выявление конфигурации ИИ-агента</strong><p>Limit public disclosure of agent tools, services, configuration, and workflows.</p></a>
<a class="relation-item" href="/techniques/AML.T0084.001/"><span class="relation-id">AML.T0084.001</span><strong>Определения инструментов</strong><p>Avoid publicly documenting sensitive agent tool definitions and capabilities.</p></a>
<a class="relation-item" href="/techniques/AML.T0084.002/"><span class="relation-id">AML.T0084.002</span><strong>Триггеры активации</strong><p>Avoid publicly disclosing agent activation keywords, events, and workflows.</p></a>
<a class="relation-item" href="/techniques/AML.T0084.003/"><span class="relation-id">AML.T0084.003</span><strong>Цепочки вызовов</strong><p>Limit public disclosure of agent call chains and execution-sink details.</p></a>
<a class="relation-item" href="/techniques/AML.T0095/"><span class="relation-id">AML.T0095</span><strong>Поиск на открытых сайтах и доменах</strong><p>Limit public technical and organizational information that reveals the AI stack, services, personnel, or other targeting details on websites and domains.</p></a>
</div>
