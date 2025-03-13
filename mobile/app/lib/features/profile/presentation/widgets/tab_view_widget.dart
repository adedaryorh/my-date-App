import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class TabViewWidget extends StatelessWidget {
  const TabViewWidget({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 2,
      child: Expanded(
        child: Scaffold(
          appBar: AppBar(
            automaticallyImplyLeading: false,
            bottom: TabBar(
              dividerColor: Colors.transparent,
              indicatorWeight: 1,
              indicatorSize: TabBarIndicatorSize.tab,
              indicatorPadding: const EdgeInsets.symmetric(horizontal: 20),
              unselectedLabelStyle: context.textTheme.bodyLarge
                  ?.copyWith(fontWeight: FontWeight.w600),
              labelStyle: context.textTheme.bodyLarge?.copyWith(
                fontWeight: FontWeight.w600,
                color: context.colorScheme.primary,
              ),
              tabs: const [
                Tab(
                  text: 'Contents',
                ),
                Tab(text: 'Celebrations'),
              ],
            ),
          ),
          body: const TabBarView(
            children: [
              Center(child: Text('No Content')),
              Center(child: Text('No Celebrations')),
            ],
          ),
        ),
      ),
    );
  }
}
