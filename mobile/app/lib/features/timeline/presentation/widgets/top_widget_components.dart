import 'package:celebut/core/core.dart';
import 'package:celebut/features/timeline/presentation/widgets/report_dialog.dart';
import 'package:celebut/features/timeline/presentation/widgets/user_avatar.dart';
import 'package:flutter/material.dart';

class TopWidgetComponents extends StatelessWidget {
  const TopWidgetComponents({
    required this.name,
    required this.time,
    super.key,
  });

  final String name;
  final String time;

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Row(
          children: [
            const UserAvatar(),
            const Space(20),
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  name,
                  style: context.textTheme.bodyLarge,
                ),
                Text(
                  time,
                  style: context.textTheme.bodySmall?.copyWith(
                    color: const Color(0xff96A7AF),
                  ),
                ),
              ],
            ),
          ],
        ),
        PopupMenuButton<int>(
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(100),
          ),
          elevation: 10,
          itemBuilder: (context) => [
            PopupMenuItem<int>(
              value: 0,
              child: Row(
                children: [
                  Image.asset(
                    AppAssets.reportFlag,
                    width: 15,
                    height: 15,
                  ),
                  const SizedBox(width: 8),
                  const Text('Report this Content'),
                ],
              ),
            ),
          ],
          onSelected: (value) {
            if (value == 0) {
              showDialog<void>(
                context: context,
                builder: (BuildContext context) {
                  return const ReportDialog();
                },
              );
            }
          },
          child: const Icon(
            Icons.more_vert,
            size: 21,
          ),
        ),
      ],
    );
  }
}
