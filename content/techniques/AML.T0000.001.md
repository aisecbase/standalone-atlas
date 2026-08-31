---
atlas_id: AML.T0000.001
atlas_type: technique
attack_ref_id: T1596
attack_ref_url: https://attack.mitre.org/techniques/T1596/
created_date: "2021-05-13"
description: Репозитории препринтов, такие как arXiv, содержат новейшие научные статьи, которые ещё не прошли рецензирование. В них могут храниться исследовательские заметки или технические отчёты, которые обычно не публикуются в...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Pre-Print Repositories
subtechnique_count: 0
subtechnique_of: AML.T0000
tactics:
    - AML.TA0002
title: Репозитории препринтов
url: /techniques/AML.T0000.001/
---

Репозитории препринтов, такие как arXiv, содержат новейшие научные статьи, которые ещё не прошли рецензирование.

В них могут храниться исследовательские заметки или технические отчёты, которые обычно не публикуются в журналах или материалах конференций.

Репозитории препринтов также служат централизованной площадкой для распространения статей, принятых к публикации в журналах.

Поиск в репозиториях препринтов позволяет злоумышленникам получить относительно актуальное представление о том, над чем работают исследователи в организации-жертве.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000/"><span class="relation-id">AML.T0000</span><strong>Поиск в открытых технических базах данных</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000.000/"><span class="relation-id">AML.T0000.000</span><strong>Журналы и материалы конференций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.002/"><span class="relation-id">AML.T0000.002</span><strong>Технические блоги</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0002 Разведка</span><p>Мы определили подход к обнаружению вредоносных URL на основе машинного обучения как репрезентативный подход и потенциальную цель по статье [URLNet: Learning a URL representation with deep learning for malicious URL detection](https://arxiv.org/abs/1802.03162), найденной на arXiv (репозитории препринтов).</p></a>
</div>
