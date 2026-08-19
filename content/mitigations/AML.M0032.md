---
atlas_id: AML.M0032
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - Cyber
created_date: "2025-11-25"
description: Определяйте границы безопасности вокруг агентных инструментов и источников данных с помощью таких методов, как доступ через API, изоляция контейнеров, изоляция выполнения кода в песочнице и ограничение частоты вызова...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
    - Deployment
modified_date: "2026-07-31"
source_name: Segmentation of AI Agent Components
technique_count: 7
title: Сегментация компонентов ИИ-агента
url: /mitigations/AML.M0032/
---

Определяйте границы безопасности вокруг агентных инструментов и источников данных с помощью таких методов, как доступ через API, изоляция контейнеров, изоляция выполнения кода в песочнице и ограничение частоты вызова инструментов. При изоляции в песочнице ограничивайте доступ к ресурсам и сети, а перед каждым запуском собирайте контейнер или виртуальную машину из чистого базового образа. Это ограничивает распространение недоверенных процессов или возможной компрометации по системе.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006/"><span class="relation-id">AML.T0006</span><strong>Активное сканирование</strong><p>Segment AI agent components so an exposed service does not reveal or provide reachability to additional internal components.</p></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для выполнения небезопасных действий, влияющих на другие компоненты.</p></a>
<a class="relation-item" href="/techniques/AML.T0085/"><span class="relation-id">AML.T0085</span><strong>Данные из ИИ-сервисов</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора чувствительных данных из ИИ-сервисов.</p></a>
<a class="relation-item" href="/techniques/AML.T0085.000/"><span class="relation-id">AML.T0085.000</span><strong>Базы данных RAG</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора чувствительных данных из баз данных RAG.</p></a>
<a class="relation-item" href="/techniques/AML.T0085.001/"><span class="relation-id">AML.T0085.001</span><strong>Инструменты ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора чувствительных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0086/"><span class="relation-id">AML.T0086</span><strong>Эксфильтрация через вызов инструмента ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для компрометации источников чувствительных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0098/"><span class="relation-id">AML.T0098</span><strong>Сбор учетных данных через инструменты ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора учетных данных.</p></a>
</div>
