# Changelog

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.1.4...plugins-destination-neo4j-v2.0.0) (2023-04-18)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([97d4a9b](https://github.com/cloudquery/cloudquery/commit/97d4a9bee95ece4840bce7e2b9b8f2b23b3fa035))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [1.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.1.3...plugins-destination-neo4j-v1.1.4) (2023-04-04)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [1.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.1.2...plugins-destination-neo4j-v1.1.3) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.1.1...plugins-destination-neo4j-v1.1.2) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))
* **neo4j:** Logging and small fixes ([#8825](https://github.com/cloudquery/cloudquery/issues/8825)) ([9831de4](https://github.com/cloudquery/cloudquery/commit/9831de48b52570c2df39f0583cc4557f23a13ce5))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.1.0...plugins-destination-neo4j-v1.1.1) (2023-03-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.8...plugins-destination-neo4j-v1.1.0) (2023-02-28)


### Features

* **deps:** Update Neo4j plugin-sdk to v1.39.1 ([#8330](https://github.com/cloudquery/cloudquery/issues/8330)) ([49715d9](https://github.com/cloudquery/cloudquery/commit/49715d9b8850b62dfea21a22cbf32dd34d2b37bb))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [1.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.7...plugins-destination-neo4j-v1.0.8) (2023-02-21)


### Bug Fixes

* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## [1.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.6...plugins-destination-neo4j-v1.0.7) (2023-02-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.5...plugins-destination-neo4j-v1.0.6) (2023-02-07)


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.4...plugins-destination-neo4j-v1.0.5) (2023-01-31)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.3...plugins-destination-neo4j-v1.0.4) (2023-01-24)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.2...plugins-destination-neo4j-v1.0.3) (2023-01-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.27.0 ([#6856](https://github.com/cloudquery/cloudquery/issues/6856)) ([545799b](https://github.com/cloudquery/cloudquery/commit/545799bb0481087e187b5f27c88f5dde9c99f2f0))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.1...plugins-destination-neo4j-v1.0.2) (2023-01-10)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.18.0 ([#6339](https://github.com/cloudquery/cloudquery/issues/6339)) ([158365a](https://github.com/cloudquery/cloudquery/commit/158365a78dfa4389074f716a0f581f18fedc1080))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.19.0 ([#6363](https://github.com/cloudquery/cloudquery/issues/6363)) ([ae6967c](https://github.com/cloudquery/cloudquery/commit/ae6967c22002c554a083f444eb611ac3e6d2698f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([#6553](https://github.com/cloudquery/cloudquery/issues/6553)) ([392b848](https://github.com/cloudquery/cloudquery/commit/392b848b3124f9cf28f6234fdb9a43d671069879))
* **deps:** Update plugin-sdk to v1.21.0 for destinations ([#6419](https://github.com/cloudquery/cloudquery/issues/6419)) ([f3b989f](https://github.com/cloudquery/cloudquery/commit/f3b989f7cbe335481dc01ad2a56cf7eff48e01d5))
* **destinations:** Handle nulls in JSONs ([#6466](https://github.com/cloudquery/cloudquery/issues/6466)) ([f434f00](https://github.com/cloudquery/cloudquery/commit/f434f00285ed8fc5edbacf03194fb983d4d98f86))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-neo4j-v1.0.0...plugins-destination-neo4j-v1.0.1) (2023-01-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6214](https://github.com/cloudquery/cloudquery/issues/6214)) ([53b2415](https://github.com/cloudquery/cloudquery/commit/53b241508d7511d4b5fa74cc4262d180c1e6df66))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))

## 1.0.0 (2023-01-01)


### Features

* Add Neo4j destination ([#6162](https://github.com/cloudquery/cloudquery/issues/6162)) ([0f4398b](https://github.com/cloudquery/cloudquery/commit/0f4398b7070aeabb2a7b595487551dd8d21fe6ec))

## Changelog
