import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class IndustryPicker extends StatefulWidget {
  const IndustryPicker({
    super.key,
  });

  @override
  State<IndustryPicker> createState() => _IndustryPickerState();
}

class _IndustryPickerState extends State<IndustryPicker> {
  String? selectedOption;

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
        height: 300,
        width: 300,
        child: ListView(
          children: [
            for (int index = 0;
                index < AppConstants.industryTypes.length;
                index++)
              CustomRadioListTile(
                option: AppConstants.industryTypes[index],
                selectedOption: selectedOption,
                onSelected: (String value) {
                  selectedOption = value;
                  setState(() {});
                },
              ),
          ],
        ),
      ),
    );
  }
}
