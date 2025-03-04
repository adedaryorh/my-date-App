import 'package:celebut/apps/shared/settings/presentation/widgets/language_option_item.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class SelectLanguage extends StatefulWidget {
  const SelectLanguage({
    super.key,
  });

  @override
  State<SelectLanguage> createState() => _SelectLanguageState();
}

class _SelectLanguageState extends State<SelectLanguage> {
  String? selectLanguage;
  @override
  Widget build(BuildContext context) {
    final flags = AppConstants.languageOptions.keys.toList();
    final languages = AppConstants.languageOptions.values.toList();
    return Scaffold(
      body: Container(
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 20),
        child: Stack(
          children: [
            ListView(
              shrinkWrap: true,
              children: [
                const SizedBox(
                  height: 20,
                ),
                Text(
                  'Select Language',
                  style: context.textTheme.bodyLarge,
                ),
                const SizedBox(
                  height: 20,
                ),
                ...List.generate(
                  flags.length,
                  (index) {
                    return LanguageOptionItem(
                      text: languages[index],
                      tapped: () {
                        setState(() {
                          selectLanguage = languages[index];
                        });
                      },
                      flag: flags[index],
                      isSelected: selectLanguage == languages[index],
                    );
                  },
                ),
              ],
            ),
            IgnorePointer(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Container(
                    height: 3,
                    width: 100,
                    color: const Color(0xff14202D).withOpacity(0.13),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: MainButton(
              loading: false,
              text: 'Save',
              pressed: () {},
            ),
          ),
          const Space(30),
        ],
      ),
    );
  }
}
