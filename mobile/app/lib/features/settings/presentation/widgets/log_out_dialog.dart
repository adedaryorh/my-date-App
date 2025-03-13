import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class LogOutDialog extends StatelessWidget {
  const LogOutDialog({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Dialog(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(15),
      ),
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 20),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Row(
              children: [
                const Spacer(),
                Text(
                  'Sign Out',
                  style: context.textTheme.titleMedium,
                ),
                const Spacer(),
                InkWell(
                  onTap: () => Navigator.of(context).pop(),
                  child: Container(
                    height: 36,
                    width: 36,
                    decoration: BoxDecoration(
                      borderRadius: const BorderRadius.all(
                        Radius.circular(8),
                      ),
                      border: Border.all(color: const Color(0xffD6D6D6)),
                    ),
                    child: const Icon(
                      Icons.close,
                      color: Color(0xff101010),
                      size: 15,
                    ),
                  ),
                ),
              ],
            ),
            const Space(30),
            const Text('Do you want to sign out?'),
            const Space(30),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Container(
                  alignment: Alignment.center,
                  padding: const EdgeInsets.symmetric(horizontal: 50),
                  height: 48,
                  decoration: BoxDecoration(
                    borderRadius: const BorderRadius.all(
                      Radius.circular(50),
                    ),
                    border: Border.all(color: const Color(0xffD6D6D6)),
                  ),
                  child: Text(
                    'No',
                    style: context.textTheme.bodyLarge,
                  ),
                ),
                const Space(20),
                Container(
                  padding: const EdgeInsets.symmetric(horizontal: 50),
                  alignment: Alignment.center,
                  height: 48,
                  decoration: BoxDecoration(
                    color: context.colorScheme.primary,
                    borderRadius: const BorderRadius.all(
                      Radius.circular(50),
                    ),
                  ),
                  child: Text(
                    'Yes',
                    style: context.textTheme.bodyLarge,
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
