import 'package:celebut/apps/shared/settings/presentation/widgets/log_out_dialog.dart';
import 'package:celebut/apps/shared/settings/presentation/widgets/select_language.dart';
import 'package:flutter/material.dart';

extension XBuildContext<T> on BuildContext {
  ThemeData get theme => Theme.of(this);
  TextTheme get textTheme => theme.textTheme;
  ColorScheme get colorScheme => theme.colorScheme;
  Color? get textColor => theme.textTheme.bodyMedium?.color;
  Size get screenSize => MediaQuery.sizeOf(this);
  EdgeInsets get edgeInset => MediaQuery.viewPaddingOf(this);
  RoundedRectangleBorder get modalShape => const RoundedRectangleBorder(
        borderRadius: BorderRadius.only(
          topLeft: Radius.circular(16),
          topRight: Radius.circular(16),
        ),
      );

  Future<void> showLanguageOptionsSheet() async {
    await showModalBottomSheet<void>(
      context: this,
      isScrollControlled: true,
      shape: modalShape,
      builder: (context) => SizedBox(
        height: context.screenSize.height * 0.6,
        child: const SelectLanguage(),
      ),
    );
    return;
  }

  Future<void> showLogOutDialog() {
    return showDialog<void>(
      context: this,
      builder: (BuildContext context) {
        return const LogOutDialog();
      },
    );
  }
}
