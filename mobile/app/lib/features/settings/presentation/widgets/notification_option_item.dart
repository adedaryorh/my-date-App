import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class NotificationOptionItem extends StatefulWidget {
  const NotificationOptionItem({
    required this.title,
    super.key,
  });

  final String title;

  @override
  State<NotificationOptionItem> createState() => _NotificationOptionItemState();
}

class _NotificationOptionItemState extends State<NotificationOptionItem> {
  bool isSwitched = false;
  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Text(
          widget.title,
          style: context.textTheme.bodyMedium
              ?.copyWith(fontWeight: FontWeight.w400),
        ),
        Transform.scale(
          scale: 0.7,
          child: Switch(
            value: isSwitched,
            onChanged: (value) {
              setState(() {
                isSwitched = value;
              });
            },
            inactiveThumbColor: Colors.grey,
          ),
        ),
      ],
    );
  }
}
