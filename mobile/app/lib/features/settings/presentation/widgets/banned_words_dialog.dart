import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class BannedWordsDialog extends StatefulWidget {
  const BannedWordsDialog({super.key});

  @override
  State<BannedWordsDialog> createState() => _BannedWordsDialogState();
}

class _BannedWordsDialogState extends State<BannedWordsDialog> {
  final TextEditingController bannedWordCtr = TextEditingController();
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
                  'Add Banned Word',
                  style: context.textTheme.titleSmall,
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
            const Space(20),
            TextFormInput(
              controller: bannedWordCtr,
              labelText: 'Enter a word you want to ban',
            ),
            const Space(20),
            Container(
              alignment: Alignment.center,
              padding: const EdgeInsets.symmetric(horizontal: 50),
              height: 48,
              decoration: BoxDecoration(
                borderRadius: const BorderRadius.all(
                  Radius.circular(50),
                ),
                color: context.colorScheme.primary,
              ),
              child: Text(
                'Add',
                style: context.textTheme.bodyLarge,
              ),
            ),
            const Space(20),
          ],
        ),
      ),
    );
  }
}
