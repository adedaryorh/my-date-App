import 'dart:io';

import 'package:celebut/core/core.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

final class DialogAction {
  final String text;
  final void Function()? onPressed;
  final TextStyle? textStyle;
  final bool isDefaultAction;

  DialogAction({
    required this.text,
    required this.onPressed,
    this.textStyle,
    this.isDefaultAction = false,
  });
}

class PlatformAlertDialog extends StatelessWidget {
  final String? title;
  final Widget content;
  final List<DialogAction> actions;

  const PlatformAlertDialog({
    required this.content,
    required this.actions,
    super.key,
    this.title,
  });

  @override
  Widget build(BuildContext context) {
    return Platform.isAndroid
        ? AlertDialog(
            title: title != null
                ? Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        title!,
                        style: context.textTheme.bodyLarge,
                      ),
                      // Image.asset(
                      //   AppAssets.logo,
                      //   height: 23,
                      //   width: 96,
                      // ),
                    ],
                  )
                : null,
            content: content,
            backgroundColor: Colors.white,
            shape:
                RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
            actions: actions
                .map(
                  (e) => TextButton(
                    onPressed: e.onPressed,
                    child: Text(
                      e.text,
                      style: e.textStyle,
                    ),
                  ),
                )
                .toList(),
          )
        : CupertinoAlertDialog(
            title: title != null
                ? Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        title!,
                        style: context.textTheme.bodyLarge,
                      ),
                      // Image.asset(
                      //   AppAssets.logo,
                      //   height: 23,
                      //   width: 96,
                      // ),
                    ],
                  )
                : null,
            content: content,
            actions: actions
                .map(
                  (e) => CupertinoDialogAction(
                    isDefaultAction: e.isDefaultAction,
                    onPressed: e.onPressed,
                    child: Text(
                      e.text,
                      style: e.textStyle,
                    ),
                  ),
                )
                .toList(),
          );
  }
}
