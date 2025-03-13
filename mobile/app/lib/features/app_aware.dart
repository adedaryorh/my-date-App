import 'package:celebut/core/core.dart';
import 'package:celebut/core/services/dialog_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:overlay_support/overlay_support.dart';

class AppAware extends ConsumerStatefulWidget {
  final Widget child;

  const AppAware({
    super.key,
    required this.child,
  });
  @override
  AppAwareState createState() => AppAwareState();
}

class AppAwareState extends ConsumerState<AppAware> {
  OverlaySupportEntry? _snackBarEntry;
  OverlaySupportEntry? _dialogEntry;
  @override
  void initState() {
    super.initState();
    final dialog = ref.read(dialogServiceProvider);
    dialog.stream.listen((event) {
      switch (event.displayType) {
        case DisplayType.snackbar:
          _snackBarEntry?.dismiss();
          _snackBarEntry = showOverlayNotification(
            (context) {
              return CustomSnackBar(
                title: event.title,
                message: event.message,
                background: event.status == Status.success
                    ? Colors.green
                    : Theme.of(context).colorScheme.error,
                onDismiss: _snackBarEntry!.dismiss,
              );
            },
            position: event.position,
            duration: event.duration,
          );
          HapticFeedback.mediumImpact();
        case DisplayType.dialog:
          _dialogEntry?.dismiss();
          _dialogEntry = showOverlay(
            (context, _) {
              return Material(
                color: Colors.black38,
                child: PlatformAlertDialog(
                  title: event.title,
                  content: Text(
                    event.message,
                    style: context.textTheme.bodyMedium,
                  ),
                  actions: [
                    if (event.dismissible)
                      DialogAction(
                        text: 'Close',
                        onPressed: () => _dialogEntry!.dismiss(),
                      ),
                    if (event.action != null)
                      DialogAction(
                        isDefaultAction: true,
                        text: event.action!.text,
                        onPressed: () {
                          event.action!.onPressed?.call();
                          _dialogEntry?.dismiss();
                        },
                      ),
                  ],
                ),
              );
            },
            duration: Duration.zero,
          );
          HapticFeedback.heavyImpact();
      }
    });
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return OverlaySupport(
      child: widget.child,
    );
  }
}
