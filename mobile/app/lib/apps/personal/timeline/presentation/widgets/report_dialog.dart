import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ReportDialog extends StatefulWidget {
  const ReportDialog({
    super.key,
  });

  @override
  State<ReportDialog> createState() => _ReportDialogState();
}

class _ReportDialogState extends State<ReportDialog> {
  String? selectedOption;
  final TextEditingController reportCtr = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Dialog(
      backgroundColor: Colors.white,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(15),
      ),
      child: Container(
        padding: const EdgeInsets.symmetric(
          horizontal: 20,
          vertical: 20,
        ),
        child: ListView(
          shrinkWrap: true,
          children: [
            Row(
              children: [
                IconButton(
                  onPressed: () {
                    context.pop();
                  },
                  icon: Icon(
                    Icons.arrow_back_ios,
                    size: 15,
                    color: context.colorScheme.primary,
                  ),
                ),
                const SizedBox(
                  width: 20,
                ),
                Text(
                  'Report an issue',
                  style: context.textTheme.titleMedium?.copyWith(
                    color: context.colorScheme.primary,
                  ),
                ),
              ],
            ),
            const SizedBox(
              height: 30,
            ),
            const Text(
              'Help us understand the problem. '
              'What is going on with this Celebration?',
            ),
            const SizedBox(
              height: 20,
            ),
            for (int index = 0;
                index < AppConstants.reportOptions.length;
                index++)
              CustomRadioListTile(
                option: AppConstants.reportOptions[index],
                selectedOption: selectedOption,
                onSelected: (String value) {
                  selectedOption = value;
                  setState(() {});
                },
              ),
            const SizedBox(
              height: 10,
            ),
            const Text(
              'Other',
            ),
            TextFormInput(
              controller: reportCtr,
              decoration: InputDecoration(
                hintText: 'Report issue not listed above',
                hintStyle: context.textTheme.bodyMedium?.copyWith(
                  fontWeight: FontWeight.w400,
                  color: const Color(0xffBDBDBD),
                ),
                isDense: true,
                border: OutlineInputBorder(
                  borderSide: const BorderSide(color: Color(0xffF3F3F3)),
                  borderRadius: BorderRadius.circular(7),
                ),
                focusedBorder: OutlineInputBorder(
                  borderSide: const BorderSide(color: Color(0xffF3F3F3)),
                  borderRadius: BorderRadius.circular(7),
                ),
                enabledBorder: OutlineInputBorder(
                  borderSide: const BorderSide(color: Color(0xffF3F3F3)),
                  borderRadius: BorderRadius.circular(7),
                ),
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            MainButton(
              loading: false,
              text: 'Report this Content',
              pressed: () {},
            ),
            const SizedBox(
              height: 50,
            ),
            Center(
              child: RichText(
                selectionColor: context.colorScheme.primary,
                text: TextSpan(
                  text: 'Learn more',
                  style: context.textTheme.bodySmall?.copyWith(
                    color: context.colorScheme.primary,
                  ),
                  children: <TextSpan>[
                    TextSpan(
                      text: ' about reporting violations of our rules',
                      style: context.textTheme.bodySmall,
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
