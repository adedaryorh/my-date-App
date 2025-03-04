import 'package:celebut/apps/apps.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

final pRouteStateProvider = ChangeNotifierProvider((ref) {
  return PersistentTabController();
});

class MainEntryPersonalView extends ConsumerStatefulWidget {
  const MainEntryPersonalView({required this.child, super.key});

  final Widget child;

  @override
  ConsumerState<MainEntryPersonalView> createState() =>
      _MainEntryPersonalViewState();
}

class _MainEntryPersonalViewState extends ConsumerState<MainEntryPersonalView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          widget.child,
          Align(
            alignment: Alignment.bottomCenter,
            child: Container(
              padding: const EdgeInsets.all(19),
              margin: const EdgeInsets.only(right: 24, left: 24, bottom: 30),
              decoration: const BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.all(Radius.circular(10)),
                boxShadow: [
                  BoxShadow(
                    color: Colors.black26,
                    offset: Offset(0, 20),
                    blurRadius: 20,
                  ),
                ],
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  InkWell(
                    onTap: () => _onTap(0),
                    child: Image.asset(
                      AppAssets.bottomNavCelebrate,
                      height: 41,
                      width: 41,
                    ),
                  ),
                  InkWell(
                    onTap: () => _onTap(1),
                    child: Image.asset(
                      AppAssets.bottomNavAvatar,
                      height: 42,
                      width: 41,
                    ),
                  ),
                  InkWell(
                    onTap: () => _onTap(2),
                    child: Image.asset(
                      AppAssets.timeline,
                      height: 31,
                      width: 31,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }

  void _onTap(int value) {
    ref.read(pRouteStateProvider).index = value;
    switch (value) {
      case 0:
        context.goNamed(AppRoute.pCreateCelebration.name);
      case 1:
        context.goNamed(AppRoute.pCelebration.name);
      case 2:
        context.goNamed(AppRoute.pTimeline.name);
    }
  }
}
