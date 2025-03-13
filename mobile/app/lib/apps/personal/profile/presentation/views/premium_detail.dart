import 'package:celebut/apps/personal/profile/presentation/widgets/title_widget.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class PremiumDetail extends StatefulWidget {
  const PremiumDetail({super.key});

  @override
  State<PremiumDetail> createState() => _PremiumDetailState();
}

class _PremiumDetailState extends State<PremiumDetail> {
  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints constraints) {
        return Scaffold(
          body: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              TitleWidget(
                constraints: constraints,
              ),
              const Space(20),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 20),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Basic',
                      style: context.textTheme.bodyLarge
                          ?.copyWith(fontWeight: FontWeight.w500),
                    ),
                    const Space(10),
                    RichText(
                      selectionColor: context.colorScheme.primary,
                      text: TextSpan(
                        text: '£4.5',
                        style: context.textTheme.displaySmall
                            ?.copyWith(fontWeight: FontWeight.w600),
                        children: <TextSpan>[
                          TextSpan(
                            text: ' /month',
                            style: context.textTheme.bodyMedium?.copyWith(
                              fontWeight: FontWeight.w500,
                            ),
                          ),
                        ],
                      ),
                    ),
                    const Space(50),
                    Text(
                      'Features:',
                      style: context.textTheme.bodyLarge
                          ?.copyWith(fontWeight: FontWeight.w500),
                    ),
                    const Space(20),
                    ...List.generate(
                      AppConstants.featureList.length,
                      (index) {
                        return Padding(
                          padding: const EdgeInsets.only(bottom: 15),
                          child: Row(
                            children: [
                              const Icon(
                                Icons.check,
                                size: 25,
                              ),
                              const Space(15),
                              Flexible(
                                child: Text(
                                  AppConstants.featureList[index],
                                ),
                              )
                            ],
                          ),
                        );
                      },
                    ),
                  ],
                ),
              ),
            ],
          ),
          bottomNavigationBar: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 20),
                child: MainButton(
                  loading: false,
                  text: 'Subscribe',
                  pressed: () {},
                ),
              ),
              const Space(30),
            ],
          ),
        );
      },
    );
  }
}
