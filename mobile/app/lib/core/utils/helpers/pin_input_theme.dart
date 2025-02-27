import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:pinput/pinput.dart';

PinTheme buildPinTheme(
  BuildContext context, [
  Color borderColor = Colors.transparent,
]) {
  return PinTheme(
    width: 64,
    height: 64,
    margin: const EdgeInsets.only(left: 5),
    textStyle: context.textTheme.bodyMedium,
    decoration: BoxDecoration(
      color: const Color(0xffEDEDED),
      borderRadius: const BorderRadius.all(Radius.circular(10)),
      border: Border.all(
        color: borderColor,
      ),
    ),
  );
}
