import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class TextFormInput extends StatelessWidget {
  const TextFormInput(
      {required this.controller,
      super.key,
      this.autoValidateMode,
      this.validator,
      this.obscureText = false,
      this.suffixIcon,
      this.onChanged,
      this.onTap,
      this.onFieldSubmitted,
      this.onEditingComplete,
      this.enabled,
      this.fillColor = Colors.white,
      this.filled = false,
      this.readOnly = false,
      this.focusNode,
      this.keyboardType,
      this.errorText,
      this.textInputAction = TextInputAction.done,
      this.labelText,
      this.prefixIcon,
      this.maxLines,
      this.decoration,
      this.labelStyle,
      this.inputFormatters});

  final bool obscureText;
  final Widget? suffixIcon;
  final bool? enabled;
  final bool readOnly;
  final bool filled;
  final Color fillColor;
  final TextInputType? keyboardType;
  final TextEditingController controller;
  final TextInputAction? textInputAction;
  final AutovalidateMode? autoValidateMode;
  final String? Function(String?)? validator;
  final void Function()? onEditingComplete;
  final void Function(String)? onChanged;
  final void Function(String)? onFieldSubmitted;
  final void Function()? onTap;
  final FocusNode? focusNode;
  final String? errorText;
  final String? labelText;
  final Widget? prefixIcon;
  final int? maxLines;
  final InputDecoration? decoration;
  final List<TextInputFormatter>? inputFormatters;
  final TextStyle? labelStyle;

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      inputFormatters: inputFormatters,
      maxLines: maxLines,
      enabled: enabled,
      controller: controller,
      readOnly: readOnly,
      focusNode: focusNode,
      autovalidateMode: autoValidateMode,
      textCapitalization: TextCapitalization.words,
      obscureText: obscureText,
      obscuringCharacter: '*',
      style: Theme.of(context).textTheme.bodyMedium,
      textInputAction: textInputAction,
      keyboardType: keyboardType,
      onChanged: onChanged,
      onEditingComplete: () => FocusScope.of(context).nextFocus(),
      onFieldSubmitted: onFieldSubmitted,
      validator: validator,
      onTap: onTap,
      decoration: InputDecoration(
        labelText: labelText,
        labelStyle: labelStyle ??
            context.textTheme.bodyMedium
                ?.copyWith(fontWeight: FontWeight.w600, color: Colors.black45),
        prefixIcon: prefixIcon,
        fillColor: Colors.white,
        suffixIcon: suffixIcon,
      ),
    );
  }
}
