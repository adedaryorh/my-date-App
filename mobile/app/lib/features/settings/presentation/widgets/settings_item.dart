import 'package:flutter/material.dart';

class SettingItem extends StatelessWidget {
  const SettingItem({
    required this.title,
    super.key,
  });

  final String title;

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(bottom: 10),
      padding: const EdgeInsets.all(15),
      height: 56,
      width: double.maxFinite,
      decoration: const BoxDecoration(
        color: Color(0xffF4F4F4),
        borderRadius: BorderRadius.all(
          Radius.circular(10),
        ),
      ),
      child: Text(
        title,
      ),
    );
  }
}
