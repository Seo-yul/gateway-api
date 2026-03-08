---
description: >
  We are excited to announce the v0.5.0 release of Gateway API. For the first
  time, several of our most important Gateway API resources are graduating to
  beta. Additional, we are starting a new initiative to explore how Gateway API
  can be used for mesh and introducing new experimental concepts such as URL
  rewrites.
---

# Gateway API 베타 졸업

<small style="position:relative; top:-30px;">
  :octicons-calendar-24: July 13, 2022 ·
  :octicons-clock-24: 5 min read
</small>

Gateway API v0.5.0 릴리스를 발표하게 되어 기쁘다. 처음으로 가장 중요한
Gateway API 리소스 중 일부가 베타로 졸업한다. 추가적으로, Gateway API가
메시에 어떻게 사용될 수 있는지 탐구하는 새로운 이니셔티브를 시작하고
URL 재작성과 같은 새로운 실험적 개념을 도입한다. 이 모든 내용과 그 이상을
아래에서 다룬다.

## Gateway API란 무엇인가?

Gateway API는 [**게이트웨이(Gateway)**][gw] 리소스(기본 네트워크 게이트웨이/프록시 서버를 나타냄)를
중심으로 한 리소스 모음으로, 표현력이 풍부하고 확장 가능하며
역할 지향적인 인터페이스를 통해 강력한 Kubernetes 서비스 네트워킹을 가능하게 한다.
이 인터페이스는 많은 벤더에 의해 구현되며 광범위한 산업 지원을 받는다.

잘 알려진 [**인그레스(Ingress)**][ing] API의 후속으로 구상된
Gateway API의 이점에는 일반적으로 사용되는 많은 네트워킹 프로토콜(예: `HTTP`, `TLS`, `TCP`, `UDP`)에 대한
명시적 지원과 Transport Layer Security(TLS)에 대한 긴밀하게 통합된 지원이
포함된다(이에 국한되지 않음). 특히 `Gateway` 리소스는 구현체가
네트워크 게이트웨이의 수명 주기를 Kubernetes API로 관리할 수 있게 한다.

Gateway API의 이점에 관심이 있는 최종 사용자라면 바로 시작하여
적합한 구현체를 찾아보길 권한다. 이 릴리스 시점에 인기 있는 API
게이트웨이와 **서비스 메시(Service Mesh)**를 위한 12개 이상의 [구현체][impl]가 있으며
빠르게 탐색을 시작할 수 있는 가이드가 제공된다.

[gw]:../../api-types/gateway.md
[ing]:https://kubernetes.io/docs/reference/kubernetes-api/service-resources/ingress-v1/
[impl]:../../implementations.md

### 시작하기

Gateway API는
[Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)와
같은 공식 Kubernetes API이다.
Gateway API는 인그레스 기능의 상위 집합을 나타내며, 더 고급 개념을
가능하게 한다. 인그레스와 마찬가지로, Kubernetes에 내장된 Gateway API의
기본 구현은 없다. 대신 많은 다양한
[구현체][impl]가 있으며, 일관되고 이식 가능한 경험을 제공하면서
기본 기술 면에서 상당한 선택권을 제공한다.

[API 개념 문서][concepts]를 살펴보고 [가이드][guides]를 확인하여
API와 작동 방식에 익숙해지기를 권한다. 실제 적용할 준비가 되면
[구현체 페이지][impl]를 열고 이미 익숙한 기존 기술에 속하는
구현체 또는 **클러스터(Cluster)** 제공자가 기본으로 사용하는 구현체를
선택한다(해당되는 경우). Gateway API는 [Custom Resource Definition
(CRD)][crd] 기반 API이므로 API를 사용하려면 클러스터에
[CRD를 설치][install-crds]해야 한다.

Gateway API에 기여하는 데 관심이 있다면 환영한다!
리포지토리에서 [새 이슈를 생성][issue]하거나
[토론][disc]에 참여하길 바란다. 또한 Slack 채널과 커뮤니티 미팅 링크가 포함된
[커뮤니티 페이지][community]도 확인하길 바란다.

[crd]:https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/
[concepts]:../../concepts/api-overview.md
[guides]:/guides/
[impl]:../../implementations.md
[install-crds]:/guides/getting-started/#installing-gateway-api
[issue]:https://github.com/kubernetes-sigs/gateway-api/issues/new/choose
[disc]:https://github.com/kubernetes-sigs/gateway-api/discussions
[community]:../../contributing/index.md

## 릴리스 주요 내용

### 베타 졸업

`v0.5.0` 릴리스는 일부 핵심 API에 대해 베타 API 버전(`v1beta1`)
릴리스로의 성숙도 성장을 나타내기 때문에 특히 역사적인 릴리스이다:

- [GatewayClass](../../api-types/gatewayclass.md)
- [Gateway](../../api-types/gateway.md)
- [HTTPRoute](../../api-types/httproute.md)

이 성과는 여러 졸업 기준의 완료로 달성되었다:

- API가 [광범위하게 구현][impl]되었다.
- 적합성 테스트가 모든 리소스에 대한 기본 커버리지를 제공하며 여러 구현체가 테스트를 통과한다.
- API 표면의 대부분이 활발히 사용되고 있다.
- Kubernetes SIG Network API 리뷰어들이 베타 졸업을 승인했다.

Gateway API 버전 관리에 대한 자세한 내용은 [공식
문서](../../concepts/versioning.md)를 참조한다. 향후
릴리스에 대한 내용은 [다음 단계](#next-steps) 섹션을 확인한다.

[impl]:../../implementations.md

### 릴리스 채널

이 릴리스에서는 안정성을 유지하면서도 실험과 반복 개발을 가능하게 하는
더 나은 균형을 위해 `experimental`과 `standard` [릴리스 채널][ch]을 도입한다.

`standard` 릴리스 채널에는 다음이 포함된다:

- 베타로 졸업한 리소스
- standard로 졸업한 필드(더 이상 실험적으로 간주되지 않음)

`experimental` 릴리스 채널에는 `standard` 릴리스 채널의 모든 것에 추가로
다음이 포함된다:

- `alpha` API 리소스
- 실험적으로 간주되며 `standard` 채널로 졸업하지 않은 필드

릴리스 채널은 내부적으로 빠른 전환을 통한 반복 개발을 가능하게 하고,
외부적으로 구현자와 최종 사용자에게 기능 안정성을 나타내는 데 사용된다.

이 릴리스에서 다음 실험적 기능이 추가되었다:

- [포트 번호를 지정하여 라우트를 게이트웨이에 연결](../../geps/gep-957/index.md)
- [URL 재작성과 경로 리디렉션](../../geps/gep-726/index.md)

[ch]:../../concepts/versioning.md#release-channels

### 기타 개선 사항

`v0.5.0` 릴리스에 포함된 전체 변경 목록은
[v0.5.0 릴리스 노트](https://github.com/kubernetes-sigs/gateway-api/releases/tag/v0.5.0)를 참조한다.

## 서비스 메시를 위한 Gateway API: GAMMA 이니셔티브
일부 서비스 메시 프로젝트는 이미 [Gateway API에 대한 지원을
구현](../../implementations.md)했다. Service Mesh Interface(SMI) API와
Gateway API 간의 상당한 중첩이 SMI 커뮤니티에서 가능한 통합에 대한
[논의를 촉발](https://github.com/servicemeshinterface/smi-spec/issues/249)했다.

Cilium Service Mesh, Consul, Istio, Kuma, Linkerd, NGINX Service Mesh,
Open Service Mesh의 대표를 포함한 서비스 메시 커뮤니티가 함께 모여
메시 관리 및 운영을 위한 Gateway API에 초점을 맞춘 Gateway API
서브프로젝트 내 전용 워크스트림인 [GAMMA
이니셔티브](../../mesh/index.md)를 구성하게 되어 기쁘다.

이 그룹은 메시 및 메시 관련 사용 사례를 위한 Gateway API 명세에 대한
리소스, 추가 사항 및 수정 사항으로 구성된 [개선
제안](../../contributing/enhancement-requests.md)을 제공할 것이다.

이 작업은 [서비스 간 트래픽을 위한 Gateway API 사용
탐구](https://docs.google.com/document/d/1T_DtMQoq2tccLAtJTpo3c0ohjm25vRS35MsestSL9QU/edit#heading=h.jt37re3yi6k5)로
시작되었으며 인증 및 인가 정책과 같은 영역에서 개선이 계속될 것이다.

## 다음 단계 {#next-steps}

프로덕션 사용 사례를 위해 API를 계속 성숙시키면서, 다음 Gateway API 릴리스에서 작업할 주요 내용은 다음과 같다:

- [gRPC][grpc] 트래픽 라우팅을 위한 [GRPCRoute][gep1016]
- [라우트 위임][pr1085]
- Layer 4 API 성숙도: [TCPRoute][tcpr], [UDPRoute][udpr] 및
  [TLSRoute][tlsr]의 베타 졸업
- [GAMMA 이니셔티브](../../mesh/index.md) - 서비스 메시를 위한 Gateway API

이 목록에 참여하고 싶은 항목이 있거나, 이 목록에 없지만 로드맵에
포함되길 원하는 항목이 있다면 Kubernetes Slack의 #sig-network-gateway-api 채널이나
주간 [커뮤니티 콜](../../contributing/index.md#meetings)에 참여하길 바란다.

[gep1016]:/geps/gep-1016/
[grpc]:https://grpc.io/
[pr1085]:https://github.com/kubernetes-sigs/gateway-api/pull/1085
[tcpr]:https://github.com/kubernetes-sigs/gateway-api/blob/main/apis/v1alpha2/tcproute_types.go
[udpr]:https://github.com/kubernetes-sigs/gateway-api/blob/main/apis/v1alpha2/udproute_types.go
[tlsr]:https://github.com/kubernetes-sigs/gateway-api/blob/main/apis/v1alpha2/tlsroute_types.go
[community]:../../contributing/index.md
