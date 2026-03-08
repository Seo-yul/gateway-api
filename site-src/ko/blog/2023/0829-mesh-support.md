---
description: >
  We are excited to announce the v0.8.0 release of Gateway API, where service
  mesh support has now reached Experimental status, we've introduced CEL
  validation and a Mesh conformance profile, and more!
---

# Gateway API v0.8.0: 서비스 메시 지원 소개!

<small style="position:relative; top:-30px;">
  :octicons-calendar-24: August 29, 2023 ·
  :octicons-clock-24: 5 min read
</small>

Gateway API v0.8.0 릴리스를 발표하게 되어 매우 기쁘다! 이번 릴리스에서
Gateway API의 **서비스 메시(Service Mesh)** 지원이 [Experimental
상태][status]에 도달했으며, CEL 검증과 같은 소규모 추가 사항도 있다.
여러분의 피드백을 기다린다!

특히 Kuma 2.3+, Linkerd 2.14+, Istio 1.16+가 모두 Gateway API
서비스 메시 지원의 완전한 적합성 구현체임을 발표하게 되어 기쁘다.

## Gateway API의 서비스 메시 지원

Gateway API의 초기 초점은 항상 **인그레스(Ingress)**(남-북) 트래픽이었지만,
동일한 기본 라우팅 개념이 서비스 메시(동-서) **트래픽(Traffic)**에도 적용 가능해야 한다는 것은
처음부터 거의 명확했다. 2022년에 Gateway API 서브프로젝트는
서비스 메시 지원을 Gateway API 리소스 프레임워크에 가장 잘 맞추는 방법을 구체적으로 검토하기 위한
전용 벤더 중립 워크스트림인 [GAMMA 이니셔티브][gamma]를 시작했으며,
Gateway API 사용자가 API에 대해 이해하고 있는 모든 것을 다시 배울 필요가 없도록 했다.

지난 1년 동안 GAMMA는 서비스 메시를 위한 Gateway API 사용과 관련된 과제와
가능한 해결책을 깊이 파고들었다. 최종 결과는
수많은 시간의 사고와 토론을 집약하고 Gateway API를
서비스 메시에 사용할 수 있는 최소 실행 가능한 경로를 제공하는
소수의 [개선 제안][geps]이다.

### Gateway API를 사용할 때 메시 라우팅은 어떻게 작동하는가?

[Gateway API 메시 라우팅 문서][mesh-routing]와 [GEP-1426]에서
모든 세부 사항을 확인할 수 있지만, Gateway API v0.8.0의 간단한 요약은
HTTPRoute가 이제 **게이트웨이(Gateway)**뿐만 아니라 Service를 `parentRef`로 가질 수 있다는 것이다.
서비스 메시 사용 사례에 대한 경험이 쌓이면서 이 영역에서 향후 GEP가 나올 것으로 예상한다.
Service에 바인딩하면 서비스 메시에서 Gateway API를 사용할 수 있지만,
여전히 다루기 어려운 흥미로운 사용 사례가 몇 가지 있다.

예를 들어, 다음과 같이 HTTPRoute를 사용하여 메시에서 A-B 테스트를 수행할 수 있다:

```yaml
  apiVersion: gateway.networking.k8s.io/v1beta1
  kind: HTTPRoute
  metadata:
    name: bar-route
  spec:
    parentRefs:
    - group: ""
      kind: Service
      name: demo-app
      port: 5000
    rules:
    - matches:
      - headers:
        - type: Exact
          name: env
          value: v1
      backendRefs:
      - name: demo-app-v1
        port: 5000
    - backendRefs:
      - name: demo-app-v2
        port: 5000
```

`demo-app` Service의 포트 5000에 대한 요청 중 `env: v1` 헤더가 있는 요청은
`demo-app-v1`으로 라우팅되고, 해당 헤더가 없는 요청은
`demo-app-v2`로 라우팅된다. 이것은 인그레스 컨트롤러가 아닌 서비스 메시에 의해
처리되므로, A/B 테스트는 애플리케이션의 호출 그래프 어디에서나 발생할 수 있다.

### 이것이 진정으로 이식 가능한지 어떻게 알 수 있는가?

Gateway API는 지원하는 모든 기능에 대해 적합성 테스트에 많은 투자를 해왔으며,
메시도 예외가 아니다. GAMMA 이니셔티브가 직면한 과제 중 하나는
이러한 테스트의 대부분이 주어진 구현이 인그레스 컨트롤러를 제공한다는
개념에 강하게 결합되어 있다는 것이었다. 많은 서비스 메시는 그렇지 않으며,
GAMMA 적합 메시가 인그레스 컨트롤러도 구현하도록 요구하는 것은
실용적이지 않았다. 이로 인해 [GEP-1709]에서 논의된 대로
Gateway API _적합성 프로파일_ 작업이 다시 시작되었다.

적합성 프로파일의 기본 개념은 Gateway API의 하위 집합을 정의하고,
구현체가 어떤 하위 집합에 적합한지 선택(및 문서화)할 수 있도록 하는 것이다.
GAMMA는 [GEP-1686]에 설명된 `Mesh`라는 새 프로파일을 추가하고 있으며,
이는 GAMMA에서 정의한 메시 기능만 확인한다.
현재 Kuma 2.3+, Linkerd 2.14+, Istio 1.16+가 모두 `Mesh` 프로파일에 적합하다.

## Gateway API v0.8.0에 다른 변경 사항은?

이 릴리스는 HTTPRoute, Gateway, GatewayClass가 GA로 졸업할 예정인
다가오는 v1.0 릴리스를 위한 Gateway API 준비에 관한 것이다.
이와 관련된 두 가지 주요 변경 사항이 있다: CEL 검증과 GEP 프로세스 변경이다.

### CEL 검증

첫 번째 주요 변경 사항은 Gateway API v0.8.0이 웹훅 검증에서
CRD에 내장된 정보를 사용하는 [CEL 검증][cel]으로의 전환을 시작한다는 것이다.
이는 사용하는 Kubernetes 버전에 따라 다른 의미를 가진다:

#### Kubernetes 1.25+

CEL 검증이 완전히 지원되며, 거의 모든 검증이 CEL로 구현된다.
(유일한 예외는 헤더 수정 필터의 헤더 이름이 대소문자 구분 없는 검증만
수행할 수 있다는 것이다. 자세한 내용은 [issue 2277]에서 확인할 수 있다.)

이러한 Kubernetes 버전에서는 검증 웹훅을 사용하지 _않는_ 것을 권장한다.

#### Kubernetes 1.23 및 1.24

CEL 검증은 지원되지 않지만, Gateway API v0.8.0 CRD는 여전히
설치할 수 있다. Kubernetes 1.25+로 업그레이드하면
이 CRD에 포함된 검증이 자동으로 적용된다.

이러한 Kubernetes 버전에서는 검증 웹훅을 계속 사용하는 것을 권장한다.

#### Kubernetes 1.22 이하

Gateway API는 [Kubernetes의 최근 5개 버전][supported-versions]에 대한 지원만
보장한다. 따라서 이러한 버전은 더 이상 Gateway API에서 지원되지 않으며,
안타깝게도 CEL 검증이 포함된 CRD가 거부되므로
Gateway API v0.8.0을 설치할 수 없다.

### GEP 프로세스 변경

Gateway API v0.8.0의 두 번째 중요한 변경 사항은 [Experimental][status]
GEP 관련 프로세스를 (필요에 의해) 면밀히 검토했다는 것이다.
이러한 GEP 중 일부가 충분히 오래 지속되어 프로젝트들이 프로덕션에서
이에 의존하게 되었으며, 이는 GEP 프로세스의 일종의 결함이다.
향후 이를 방지하기 위해, [Experimental][status]에 도달하려면
GEP가 [Standard][status]가 되기 위한 졸업 기준과
졸업 기준을 충족하지 못할 경우 GEP가 삭제되는 수습 기간을
모두 포함하도록 GEP 프로세스를 변경했다.

`v0.8.0` 릴리스에 포함된 전체 변경 목록은
[v0.8.0 릴리스 노트]를 참조한다. Gateway API 버전 관리에 대한 자세한 내용은
[공식 문서][versioning docs]를 참조한다.

## Gateway API를 어떻게 시작할 수 있는가?

Gateway API는 Kubernetes에서 로드 밸런싱, 라우팅, 서비스 메시 API의
미래를 나타낸다. 이미 20개 이상의 [구현체][impl]가
사용 가능하며(인그레스 컨트롤러와 서비스 메시 모두 포함) 목록은
계속 늘어나고 있다.

Gateway API를 시작하는 데 관심이 있다면 [API 개념 문서][concepts]를 살펴보고
[가이드][guides]를 확인하여 사용해 보길 바란다. CRD 기반 API이므로
Kubernetes 1.23+ **클러스터(Cluster)**에 최신 버전을 설치할 수 있다.

Gateway API에 기여하는 데 관심이 있다면 환영한다!
리포지토리에서 [새 이슈를 생성][issue]하거나
[토론][disc]에 참여하길 바란다. 또한 Slack 채널과 커뮤니티 미팅 링크가 포함된
[커뮤니티 페이지][community]도 확인하길 바란다. 여러분을 만나기를 기대한다!!

## 추가 읽을거리:

- [GEP-1324]는 GAMMA 목표와 몇 가지 중요한 정의에 대한 개요를 제공한다.
  이 GEP는 문제 영역에 대한 논의를 위해 읽어볼 가치가 있다.
- [GEP-1426]은 서비스 메시 내에서 트래픽을 관리하기 위해
  HTTPRoute와 같은 Gateway API 라우트 리소스를 사용하는 방법을 정의한다.
- [GEP-1686]은 [GEP-1709]의 작업을 기반으로 서비스 메시가 Gateway API에
  적합하다고 선언하기 위한 _적합성 프로파일_을 정의한다.

이것들은 [Experimental][status] 패턴이지만, GAMMA 이니셔티브가
지금까지 새로운 리소스나 필드를 도입할 필요가 없었기 때문에
[`standard` 릴리스 채널][ch]에서 사용할 수 있다는 점에 유의한다.

[gamma]:../../mesh/index.md
[status]:../../geps/overview.md#gep-states
[ch]:../../concepts/versioning.md#release-channels
[cel]:https://kubernetes.io/docs/reference/using-api/cel/
[crd]:https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/
[concepts]:../../concepts/api-overview.md
[geps]:../../contributing/enhancement-requests.md
[guides]:/guides/
[impl]:../../implementations.md
[install-crds]:/guides/getting-started/#installing-gateway-api
[issue]:https://github.com/kubernetes-sigs/gateway-api/issues/new/choose
[disc]:https://github.com/kubernetes-sigs/gateway-api/discussions
[community]:../../contributing/index.md
[mesh-routing]:../../mesh/index.md
[GEP-1426]:../../geps/gep-1294/index.md
[GEP-1324]:../../geps/gep-1324/index.md
[GEP-1686]:../../geps/gep-1686/index.md
[GEP-1709]:../../geps/gep-1709/index.md
[issue 2277]:https://github.com/kubernetes-sigs/gateway-api/issues/2277
[supported-versions]:../../concepts/versioning.md#supported-versions
[v0.8.0 release notes]:https://github.com/kubernetes-sigs/gateway-api/releases/tag/v0.8.0
[versioning docs]:../../concepts/versioning.md
