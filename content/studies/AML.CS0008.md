---
actor: Researchers at Silent Break Security
atlas_id: AML.CS0008
atlas_type: case-study
case_study_type: exercise
description: 'Proof Pudding (CVE-2019-20634) — это репозиторий кода, описывающий, как исследователи ML обошли систему защиты электронной почты ProofPoint: сначала они создали модель-копию ML-модели защиты электронной почты, а затем...'
generated: true
generated_by: atlasgen
incident_date: "2019-09-09"
incident_date_granularity: Day
incident_date_raw: "2019-09-09"
procedure:
    - description: Исследователи обнаружили, что ProofPoint Email Protection оставляла выходные оценки модели в заголовках писем.
      description_line: Исследователи обнаружили, что ProofPoint Email Protection оставляла выходные оценки модели в заголовках писем.
      tactic: AML.TA0008
      tactic_name: Выявление
      technique: AML.T0063
      technique_name: Выявление выходных данных ИИ-модели
    - description: Исследователи отправили через систему множество писем, чтобы собрать выходные данные модели из заголовков.
      description_line: Исследователи отправили через систему множество писем, чтобы собрать выходные данные модели из заголовков.
      tactic: AML.TA0000
      tactic_name: Доступ к ИИ-модели
      technique: AML.T0047
      technique_name: Продукт или сервис с поддержкой ИИ
    - description: |-
        Исследователи использовали письма и собранные оценки как набор данных, на котором обучили рабочую копию модели ProofPoint.

        С помощью простой корреляции они определили, какая переменная оценки в целом отражает безопасность письма. В этом случае была выбрана переменная "mlxlogscore", поскольку она была связана со spam, phish и core mlx, и ее использовали как метку. Каждое значение "mlxlogscore" обычно находилось в диапазоне от 1 до 999: чем выше оценка, тем безопаснее образец. Обучение выполнялось с использованием искусственной нейронной сети (ANN) и токенизации Bag of Words.
      description_line: 'Исследователи использовали письма и собранные оценки как набор данных, на котором обучили рабочую копию модели ProofPoint. С помощью простой корреляции они определили, какая переменная оценки в целом отражает безопасность письма. В этом случае была выбрана переменная "mlxlogscore", поскольку она была связана со spam, phish и core mlx, и ее использовали как метку. Каждое значение "mlxlogscore" обычно находилось в диапазоне от 1 до 999: чем выше оценка, тем безопаснее образец. Обучение выполнялось с использованием искусственной нейронной сети (ANN) и токенизации Bag of Words.'
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0005.001
      technique_name: Обучение прокси-модели через репликацию
    - description: |-
        Затем исследователи ML алгоритмически нашли в этой "офлайн" прокси-модели образцы, которые помогли получить нужное представление о ее поведении и влиятельных переменных.

        Примеры образцов с хорошими оценками: "calculation", "asset" и "tyson".

        Примеры образцов с плохими оценками: "software", "99" и "unsub".
      description_line: 'Затем исследователи ML алгоритмически нашли в этой "офлайн" прокси-модели образцы, которые помогли получить нужное представление о ее поведении и влиятельных переменных. Примеры образцов с хорошими оценками: "calculation", "asset" и "tyson". Примеры образцов с плохими оценками: "software", "99" и "unsub".'
      tactic: AML.TA0001
      tactic_name: Подготовка атаки на ИИ
      technique: AML.T0043.002
      technique_name: Перенос на модель чёрного ящика
    - description: В итоге сведения, полученные с помощью "офлайн" прокси-модели, позволили исследователям создавать вредоносные письма, которые получали благоприятные оценки от реальной системы защиты электронной почты ProofPoint и тем самым обходили ее.
      description_line: В итоге сведения, полученные с помощью "офлайн" прокси-модели, позволили исследователям создавать вредоносные письма, которые получали благоприятные оценки от реальной системы защиты электронной почты ProofPoint и тем самым обходили ее.
      tactic: AML.TA0011
      tactic_name: Воздействие
      technique: AML.T0015
      technique_name: Обход ИИ-модели
procedure_count: 5
references:
    - title: National Vulnerability Database entry for CVE-2019-20634
      url: https://nvd.nist.gov/vuln/detail/CVE-2019-20634
    - title: '2019 DerbyCon presentation "42: The answer to life, the universe, and everything offensive security"'
      url: https://github.com/moohax/Talks/blob/master/slides/DerbyCon19.pdf
    - title: Proof Pudding (CVE-2019-20634) Implementation on GitHub
      url: https://github.com/moohax/Proof-Pudding
    - title: '2019 DerbyCon video presentation "42: The answer to life, the universe, and everything offensive security"'
      url: https://www.youtube.com/watch?v=CsvkYoxtexQ&ab-channel=AdrianCrenshaw
reporter: ""
source_name: ProofPoint Evasion
target: ProofPoint Email Protection System
title: Обход ProofPoint
url: /studies/AML.CS0008/
---

Proof Pudding (CVE-2019-20634) — это репозиторий кода, описывающий, как исследователи ML обошли систему защиты электронной почты ProofPoint: сначала они создали модель-копию ML-модели защиты электронной почты, а затем использовали полученные сведения, чтобы обойти рабочую систему. В частности, эти сведения позволили исследователям создавать вредоносные письма, которые получали предпочтительные оценки и оставались незамеченными системой. Каждое слово в письме оценивается численно по нескольким переменным; если общая оценка письма слишком низкая, ProofPoint выводит ошибку и помечает письмо как SPAM.
