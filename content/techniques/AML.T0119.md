---
atlas_id: AML.T0119
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Злоумышленники могут отправить, опубликовать или изменить артефакт таким образом, чтобы запустить автоматизированный пайплайн его обработки. Когда пайплайн обрабатывает артефакт, злоумышленник может с помощью...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 2
source_name: Exploit Automated Artifact Processing Pipeline
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
title: Эксплуатация автоматизированного пайплайна обработки артефактов
url: /techniques/AML.T0119/
---

Злоумышленники могут отправить, опубликовать или изменить артефакт таким образом, чтобы запустить автоматизированный пайплайн его обработки. Когда пайплайн обрабатывает артефакт, злоумышленник может с помощью содержимого или конфигурации, находящихся под его контролем, эксплуатировать слабое место в логике обработки, в результате чего обработчик действует за рамками предусмотренного поведения, используя разрешения и доступ воркера обработки.

Эксплуатация может привести к тому, что воркер обработки получит доступ к локальным или внутренним ресурсам и раскроет информацию через штатные выходные данные обработки либо интерпретирует подконтрольное злоумышленнику содержимое и выполнит код. Атака запускается автоматизированной обработкой на серверной стороне и не требует от жертвы открывать, загружать или одобрять артефакт.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Сканирование уязвимостей сокращает возможности злоумышленников эксплуатировать недостатки в пайплайнах обработки артефактов.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0004 Первичный доступ</span><p>Агенты внесли в подконтрольный им репозиторий вредоносную конфигурацию набора данных HDF5. Она содержала ссылки HDF5 на внешние файлы, заданные в качестве хранилища необработанных данных. Затем агенты запросили автоматизированную обработку этого набора данных через API Hugging Face Dataset Server.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0004 Первичный доступ</span><p>The agents submitted a malicious reference dataset configuration to the automated production pipeline. Jinja2 evaluated an attacker-controlled expression and executed arbitrary Python inside a Dataset Server conversion worker running as a pod in a production Kubernetes cluster.</p></a>
</div>
