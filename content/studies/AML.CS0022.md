---
actor: Vulcan Cyber, Lasso Security
atlas_id: AML.CS0022
atlas_type: case-study
case_study_type: exercise
description: Исследователи установили, что большие языковые модели, такие как ChatGPT, могут галлюцинировать названия фиктивных программных пакетов, которые не опубликованы в репозитории пакетов. Злоумышленник может опубликовать...
generated: true
generated_by: atlasgen
incident_date: 2024-06
incident_date_granularity: Month
incident_date_raw: "2024-06-01"
procedure:
    - description: На протяжении упражнения исследователи использовали публичный API ChatGPT.
      description_line: На протяжении упражнения исследователи использовали публичный API ChatGPT.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0040
      technique_name: Доступ к API инференса ИИ-модели
    - description: |-
        Исследователи просили ChatGPT предложить программные пакеты и выявляли среди рекомендаций галлюцинации — пакеты, которых нет в публичном репозитории.

        Например, на вопрос "how to upload a model to huggingface?" модель предложила установить пакет `huggingface-cli` командой `pip install huggingface-cli`.

        Такого пакета в PyPI не существовало; реальный CLI-инструмент Hugging Face входит в пакет `huggingface_hub`.
      description_line: Исследователи просили ChatGPT предложить программные пакеты и выявляли среди рекомендаций галлюцинации — пакеты, которых нет в публичном репозитории. Например, на вопрос "how to upload a model to huggingface?" модель предложила установить пакет `huggingface-cli` командой `pip install huggingface-cli`. Такого пакета в PyPI не существовало; реальный CLI-инструмент Hugging Face входит в пакет `huggingface_hub`.
      tactic: AML.TA0008
      tactic_name: Выявление
      technique: AML.T0062
      technique_name: Выявление галлюцинированных сущностей LLM
    - description: |-
        Злоумышленник мог загрузить вредоносный пакет под галлюцинированным именем в PyPI или другие реестры пакетов.

        На практике исследователи загрузили в PyPI пустой пакет, чтобы отслеживать скачивания.
      description_line: Злоумышленник мог загрузить вредоносный пакет под галлюцинированным именем в PyPI или другие реестры пакетов. На практике исследователи загрузили в PyPI пустой пакет, чтобы отслеживать скачивания.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0060
      technique_name: Публикация галлюцинированных сущностей
    - description: |-
        Пользователь ChatGPT или другой LLM может задать похожий вопрос, получить то же галлюцинированное имя пакета и скачать вредоносный пакет.

        Исследователи показали, что несколько LLM могут выдавать одни и те же галлюцинации, и зафиксировали более 30 000 скачиваний пакета `huggingface-cli`.
      description_line: Пользователь ChatGPT или другой LLM может задать похожий вопрос, получить то же галлюцинированное имя пакета и скачать вредоносный пакет. Исследователи показали, что несколько LLM могут выдавать одни и те же галлюцинации, и зафиксировали более 30 000 скачиваний пакета `huggingface-cli`.
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0010.001
      technique_name: ПО для ИИ
    - description: В итоге пользователь загрузит вредоносный пакет, что позволит выполнить произвольный код.
      description_line: В итоге пользователь загрузит вредоносный пакет, что позволит выполнить произвольный код.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0011.001
      technique_name: Вредоносный пакет
    - description: Это может привести к различному ущербу для конечного пользователя или организации.
      description_line: Это может привести к различному ущербу для конечного пользователя или организации.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0048.003
      technique_name: Ущерб пользователям
procedure_count: 6
references:
    - title: Vulcan18's "Can you trust ChatGPT's package recommendations?"
      url: https://vulcan.io/blog/ai-hallucinations-package-risk
    - title: 'Lasso Security Research: Diving into AI Package Hallucinations'
      url: https://www.lasso.security/blog/ai-package-hallucinations
    - title: 'AIID Incident 731: Hallucinated Software Packages with Potential Malware Downloaded Thousands of Times by Developers'
      url: https://incidentdatabase.ai/cite/731/
    - title: 'Slopsquatting: When AI Agents Hallucinate Malicious Packages'
      url: https://www.trendmicro.com/vinfo/us/security/news/cybercrime-and-digital-threats/slopsquatting-when-ai-agents-hallucinate-malicious-packages
reporter: ""
source_name: ChatGPT Package Hallucination
target: ChatGPT users
title: Галлюцинация пакетов ChatGPT
url: /studies/AML.CS0022/
---

Исследователи установили, что большие языковые модели, такие как ChatGPT, могут галлюцинировать названия фиктивных программных пакетов, которые не опубликованы в репозитории пакетов. Злоумышленник может опубликовать вредоносный пакет под сгаллюцинированным названием в репозитории пакетов. Затем пользователи той же или похожей большой языковой модели могут столкнуться с той же галлюцинацией и в итоге скачать и выполнить вредоносный пакет, что может привести к различным потенциальным последствиям.
