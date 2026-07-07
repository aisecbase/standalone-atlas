---
atlas_id: AML.T0058
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут публиковать отравленную модель в публичном месте, например в реестре моделей или репозитории кода. Отравленная модель может быть новой моделью или отравленным вариантом существующей модели с...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Publish Poisoned Models
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Публикация отравленных моделей
url: /techniques/AML.T0058/
---

Злоумышленники могут публиковать отравленную модель в публичном месте, например в реестре моделей или репозитории кода. Отравленная модель может быть новой моделью или отравленным вариантом существующей модели с открытым исходным кодом. Такая модель может попасть в систему жертвы через [компрометацию цепочки поставок ИИ](/techniques/AML.T0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0023/"><span class="relation-id">AML.M0023</span><strong>Ведомость материалов ИИ</strong><p>AI BOM может помочь пользователям выявлять недоверенные артефакты моделей.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0019/"><span class="relation-id">AML.CS0019</span><strong>PoisonGPT</strong><span class="relation-meta">Актор: Mithril Security Researchers / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи загрузили модель PoisonGPT обратно в Hugging Face под именем репозитория, похожим на имя исходной модели, но с пропущенной одной буквой.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь повторно загрузил измененную модель в репозиторий Hugging Face.</p></a>
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник загрузил модель на Hugging Face. В обоих случаях, наблюдавшихся ReversingLabs, вредоносные модели не пытались имитировать популярную легитимную модель.</p></a>
</div>
