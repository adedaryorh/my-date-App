import 'package:celebut/apps/shared/app_aware.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_portal/flutter_portal.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

class AppWidget extends StatefulHookConsumerWidget {
  const AppWidget({super.key});

  @override
  ConsumerState<AppWidget> createState() => _AppWidgetState();
}

class _AppWidgetState extends ConsumerState<AppWidget> {
  @override
  Widget build(BuildContext context) {
    final appRouter = ref.watch(goRouterProvider);
    return AppAware(
      child: Portal(
        child: MaterialApp.router(
          debugShowCheckedModeBanner: false,
          title: 'Celebut',
          theme: AppTheme.light,
          routerConfig: appRouter,
        ),
      ),
    );
  }
}
