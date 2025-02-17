import 'dart:async';

import 'package:celebut/core/core.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:overlay_support/overlay_support.dart';

final dialogServiceProvider = Provider((ref) => DialogService.instance);

final class DialogService {
  static final _singleton = DialogService._internal();
  DialogService._internal();

  static DialogService get instance => _singleton;

  final StreamController<DialogModel> _controller =
      StreamController<DialogModel>();

  Stream<DialogModel> get stream => _controller.stream;

  void displayMessage(
    String message, {
    Status status = Status.failed,
    String? title,
  }) {
    _controller.add(
      DialogModel(
        title: title,
        message: message,
        status: status,
      ),
    );
  }

  void displayDialog({
    required String title,
    required String message,
    DialogAction? action,
    bool dismissible = true,
  }) {
    _controller.add(
      DialogModel(
        title: title,
        message: message,
        action: action,
        displayType: DisplayType.dialog,
        dismissible: dismissible,
      ),
    );
  }

  void dispose() => _controller.close();
}

class DialogModel {
  final String? title;
  final String message;
  final DialogAction? action;
  final NotificationPosition position;
  final Duration duration;
  final DisplayType displayType;
  final Status status;
  final bool dismissible;

  DialogModel({
    required this.title,
    required this.message,
    this.duration = const Duration(seconds: 4),
    this.position = NotificationPosition.top,
    this.action,
    this.displayType = DisplayType.snackbar,
    this.status = Status.failed,
    this.dismissible = true,
  });
}

enum Status { success, failed }

enum DisplayType { snackbar, dialog }
